package models

type GptErrorResponse struct {
	Error ErrorModel `json:"error"`
}

type ErrorModel struct {
	Message string `json:"message"`
	Type    string `json:"type"`
	Param   string `json:"param"`
	Code    string `json:"code"`
}
