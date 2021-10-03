package response

import "net/http"

type SuccessResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewSuccess() SuccessResponse {
	return SuccessResponse{
		Code:    http.StatusOK,
		Message: "success",
	}
}

func NewCreated() SuccessResponse {
	return SuccessResponse{
		Code:    http.StatusCreated,
		Message: "created",
	}
}
