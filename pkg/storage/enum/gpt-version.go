package enum

type GptEngine string

const (
	Turbo      GptEngine = "gpt-3.5-turbo"
	Four       GptEngine = "gpt-4"
	Third      GptEngine = "gpt-3"
	NotSettled GptEngine = "not-settled"
)

func (g GptEngine) String() string {
	return string(g)
}
