package gpt

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/berkayaydmr/git-ai/pkg/clients/gpt/models"
	pkgmessages "github.com/berkayaydmr/git-ai/pkg/messages"
	"github.com/berkayaydmr/git-ai/pkg/storage/enum"
	storagemodels "github.com/berkayaydmr/git-ai/pkg/storage/models"
	"github.com/berkayaydmr/git-ai/utils"
)

const (
	UserRole      = "user"
	AssistantRole = "assistant"

	temperature = 0.25

	askUrl = "https://api.openai.com/v1/chat/completions"

	MakeModelSoftwareEngineerCharacter = "You are a software engineer who is working on a project. You are questioning about somethings. You use layout to make comment by using this layout make your work."
	AskDiffQuestion                    = "You will share your ideas and information using layout in layout variables names will be fill and its represented with {{name}}. Fill the blanks with the correct answers and make your comments."
	RepositoryAndBranchNames           = "Repository name: %s\n Branches: %s, %s"
	Timeout                            = time.Minute
)

type Client struct {
	httpClient *http.Client
}

func New() *Client {
	return &Client{httpClient: &http.Client{}}
}

func (c *Client) Ask(ctx context.Context, apiKeyModel storagemodels.Profile, messages ...string) ([]models.Choice, error) {
	if apiKeyModel.GptEngine == enum.AskEveryTime {
		engine, err := utils.GetGptEngine(false)
		if err != nil {
			return nil, err
		}

		apiKeyModel.GptEngine = engine
	}

	var gptMessages []models.Message
	for _, message := range messages {
		gptMessages = append(gptMessages, splitMessages(message, apiKeyModel.GptEngine.Limit())...)
	}

	gptAskModel := &models.GptAskModel{
		Messages:    gptMessages,
		Temperature: temperature,
	}

	gptAskModel.Model = apiKeyModel.GptEngine

	tokenCount := 0
	for _, message := range gptAskModel.Messages {
		token, err := gptAskModel.Model.Encode(message.Content)
		if err != nil {
			continue
		}
		tokenCount += token
	}

	fmt.Println("Token count:", tokenCount)

	gptAskModel.MaxTokens = tokenCount

	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(gptAskModel); err != nil {
		return nil, err
	}

	fmt.Println()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, askUrl, buf)
	if err != nil {
		return nil, err
	}

	stopChannel := make(chan struct{})
	go utils.LoadingWithDots(pkgmessages.WaitingGpt, stopChannel)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKeyModel.Key))

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	close(stopChannel)

	if resp.StatusCode != http.StatusOK {
		var errResp models.GptErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&errResp); err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("\nunexpected error from gpt: \n %s \n", errResp.Error.Message)
	}

	var chatResponse models.ChatReponse
	if err := json.NewDecoder(resp.Body).Decode(&chatResponse); err != nil {
		return nil, err
	}

	return chatResponse.Choices, nil
}

func splitMessages(text string, maxTokens int) []models.Message {
	words := strings.Fields(text)
	var chunks []string
	var chunk []string
	var tokenCount int

	for _, word := range words {
		if tokenCount+len(word) > maxTokens {
			chunks = append(chunks, strings.Join(chunk, " "))
			chunk = []string{word}
			tokenCount = len(word)
		} else {
			chunk = append(chunk, word)
			tokenCount += len(word)
		}
	}

	if len(chunk) > 0 {
		chunks = append(chunks, strings.Join(chunk, " "))
	}

	var messages []models.Message
	for _, chunk := range chunks {
		messages = append(messages, models.Message{Content: chunk, Role: UserRole})
	}

	return messages
}
