package models

type GptAskModel struct {
	Model       string     `json:"model"`
	Messages    []Message `json:"messages"`
	Temperature float64    `json:"temperature"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
