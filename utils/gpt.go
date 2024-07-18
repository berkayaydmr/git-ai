package utils

import (
	"fmt"

	"github.com/berkayaydmr/git-ai/pkg/errors"
	"github.com/berkayaydmr/git-ai/pkg/storage/enum"
	"github.com/berkayaydmr/git-ai/pkg/storage/models"
)

var gptEngines = []string{
	enum.Turbo.String(),
	enum.Four.String(),
	enum.Third.String(),
	enum.NotSettled.String(),
}

func GetGptEngine() (enum.GptEngine, error) {
	for i, v := range gptEngines {
		TypeWriterEffect(fmt.Sprintf("%d. %s", i+1, v), 50)
	}

	fmt.Print("Select on of the GPT engine:")
	var selected int
	fmt.Scanln(&selected)

	if selected < 1 || selected > len(gptEngines) {
		return "", errors.ErrInvalidSelection
	}

	return enum.GptEngine(gptEngines[selected-1]), nil
}

func GetApiKeyModelFromUser(name string) (models.ApiKey, error) {
	TypeWriterEffect("Enter the API key for the GPT engine:", Faster)
	var apiKey string
	fmt.Scanln(&apiKey)
	fmt.Println()

	engine, err := GetGptEngine()
	if err != nil {
		return models.ApiKey{}, err
	}

	return models.ApiKey{Name: name, Key: apiKey, GptVersion: engine}, nil
}
