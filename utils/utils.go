package utils

import (
	"fmt"
	"time"

	"github.com/berkayaydmr/git-ai/pkg/clients/gpt/models"
)

type Speed int

const (
	Slow   Speed = 100
	Normal Speed = 75
	Fast   Speed = 50
	Faster Speed = 25
	Light  Speed = 10
)

func TypeWriterEffect(text string, speed Speed) {

	for _, char := range text {
		fmt.Print(string(char))
		time.Sleep(time.Duration(speed) * time.Millisecond)
	}

	fmt.Println()
}

func MakeChoicesString(choices []models.Choice) string {
	var result string
	for _, choice := range choices {
		result += fmt.Sprintf("%s\n", choice.String())
	}
	return result
}