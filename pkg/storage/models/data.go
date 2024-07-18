package models

import "github.com/berkayaydmr/git-ai/pkg/storage/enum"

type Data struct {
	Init    bool     `json:"init"`
	ApiKeys []ApiKey `json:"apiKeys"`
}

type ApiKey struct {
	Name       string         `json:"name"`
	Key        string         `json:"key"`
	GptVersion enum.GptEngine `json:"gptVersion"`
}
