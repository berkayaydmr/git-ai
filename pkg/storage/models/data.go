package models

import "github.com/berkayaydmr/git-ai/pkg/storage/enum"

type Data struct {
	Init     bool      `json:"init"`
	Profiles []Profile `json:"profiles"`
}

type Profile struct {
	Name      string `json:"name"`
	Key       string `json:"key"`
	GptEngine enum.GptProfileEngine `json:"gptVersion"`
}
