package enum

import (
	"github.com/pkoukk/tiktoken-go"
)

type GptProfileEngine string

const (
	Four         GptProfileEngine = "gpt-4"
	FourO        GptProfileEngine = "gpt-4-o"
	Turbo        GptProfileEngine = "gpt-3.5-turbo"
	Third        GptProfileEngine = "gpt-3"
	AskEveryTime GptProfileEngine = "ask-every-time"
)

func (g GptProfileEngine) String() string {
	return string(g)
}

func (g GptProfileEngine) CheckMessageContentExceedTokenLimit(message string) bool {
	if len(message) > gptEngineLimits[g] {
		return true
	}

	return false
}

func (g GptProfileEngine) Limit() int {
	return gptEngineLimits[g]
}

var gptEngineLimits = map[GptProfileEngine]int{
	Four:         8192,
	Turbo:        4096,
	Third:        4096,
	AskEveryTime: 4096,
}

func (g GptProfileEngine) Encode(message string) (int, error) {
	tke, err := tiktoken.GetEncoding(gptTiktokenEncodeNames[g])
	if err != nil {
		return 0, err
	}

	encoded := tke.Encode(message, nil, nil)

	return len(encoded), nil
}

var gptTiktokenEncodeNames = map[GptProfileEngine]string{
	Four:  "o200k_base",
	Turbo: "o200k_base",
	FourO: "cl100k_base",
	Third: "r50k_base",
}
