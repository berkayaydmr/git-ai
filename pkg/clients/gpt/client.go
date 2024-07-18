package gpt

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
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

	gptApiKey = "sk-proj-vKve7ZBSTxjlse8Fd28BT3BlbkFJQRFjt0t8GNGJ0jcmxkX2"

	AskDiffQuestion = "By using the diff between the branches, fill the layout"

	Timeout = time.Minute
)

type Client struct {
	httpClient *http.Client
}

func New() *Client {
	return &Client{httpClient: &http.Client{}}
}

func (c *Client) Ask(ctx context.Context, apiKeyModel storagemodels.ApiKey, messages ...string) ([]models.Choice, error) {
	var gptMessages []models.Message
	for _, message := range messages {
		gptMessages = append(gptMessages, models.Message{Content: message, Role: UserRole})
	}

	gptAskModel := &models.GptAskModel{
		Messages:    gptMessages,
		Temperature: temperature,
	}

	if apiKeyModel.GptVersion != enum.NotSettled {
		gptAskModel.Model = apiKeyModel.GptVersion.String()
	}

	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(gptAskModel); err != nil {
		return nil, err
	}

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
		return nil, fmt.Errorf("\nunexpected status code from gpt: %d", resp.StatusCode)
	}

	var chatResponse models.ChatReponse
	if err := json.NewDecoder(resp.Body).Decode(&chatResponse); err != nil {
		return nil, err
	}

	return chatResponse.Choices, nil
}
