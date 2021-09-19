package utils

import "net/http"

type SuccessResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewSuccessMessage() SuccessResponse {
	return SuccessResponse{
		Code:    http.StatusOK,
		Message: "success",
	}
}
