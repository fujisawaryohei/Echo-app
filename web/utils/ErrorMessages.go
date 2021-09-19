package utils

import (
	"net/http"

	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Errors  []ValidationError `json:"errors"`
}

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (e *ErrorResponse) SetValidationErrors(err error) {
	for _, err := range err.(validator.ValidationErrors) {
		e.Errors = append(e.Errors, ValidationError{Field: err.Field(), Message: err.Error()})
	}
}

func NewBadRequestMessage(err error) ErrorResponse {
	var errors []ValidationError
	res := ErrorResponse{
		Code:    http.StatusBadRequest,
		Message: "Bad Request",
		Errors:  errors,
	}
	res.SetValidationErrors(err)
	return res
}

func NewNotFoundMessage(err error) ErrorResponse {
	var errors []ValidationError
	res := ErrorResponse{
		Code:    http.StatusNotFound,
		Message: "Not Found",
		Errors:  errors,
	}
	return res
}

func NewInternalServerError(err error) ErrorResponse {
	var errors []ValidationError
	res := ErrorResponse{
		Code:    http.StatusInternalServerError,
		Message: "Internal Server Error",
		Errors:  errors,
	}
	return res
}
