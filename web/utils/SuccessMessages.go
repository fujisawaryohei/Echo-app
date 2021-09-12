package utils

import "net/http"

type SuccessMsg struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

func NewCreateSuccessMessage() *SuccessMsg {
	return &SuccessMsg{
		StatusCode: http.StatusOK,
		Message:    "登録が完了しました",
	}
}

func NewDeleteSuccessMessage() *SuccessMsg {
	return &SuccessMsg{
		StatusCode: http.StatusOK,
		Message:    "削除が完了しました",
	}
}
