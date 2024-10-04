package utils

import (
	"fmt"

	"github.com/berkayaydmr/git-ai/pkg/errors"
	"github.com/berkayaydmr/git-ai/pkg/storage/enum"
	"github.com/berkayaydmr/git-ai/pkg/storage/models"
)

var gptEngines = []string{
	enum.Four.String(),
	enum.Turbo.String(),
	enum.Third.String(),
	enum.AskEveryTime.String(),
}

func GetGptEngine(askEveryTime bool) (enum.GptProfileEngine, error) {
	for i, v := range gptEngines {
		if !askEveryTime && v == enum.AskEveryTime.String() {
			continue
		}

		TypeWriterEffect(fmt.Sprintf("%d. %s", i+1, v), 50)
	}

	fmt.Print("Select on of the GPT engine:")
	var selected int
	fmt.Scanln(&selected)

	if selected < 1 || selected > len(gptEngines) {
		return "", errors.ErrInvalidSelection
	}

	return enum.GptProfileEngine(gptEngines[selected-1]), nil
}

func GetGptProfileFromUser(name string) (models.Profile, error) {
	TypeWriterEffect("Enter the API key:", Faster)
	var apiKey string
	fmt.Scanln(&apiKey)
	fmt.Println()

	engine, err := GetGptEngine(true)
	if err != nil {
		return models.Profile{}, err
	}

	return models.Profile{Name: name, Key: apiKey, GptEngine: enum.GptProfileEngine(engine.String())}, nil
}
