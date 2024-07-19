package models

import "github.com/berkayaydmr/git-ai/pkg/storage/enum"

type GptAskModel struct {
	Model       enum.GptProfileEngine `json:"model"`
	Messages    []Message             `json:"messages"`
	Temperature float64               `json:"temperature"`
	MaxTokens   int                   `json:"max_tokens"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
