package response

import (
	"net/http"

	"github.com/fujisawaryohei/blog-server/codes"
	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Code    int                     `json:"code"`
	Message string                  `json:"message"`
	Errors  []codes.ValidationError `json:"errors"`
}

func (e *ErrorResponse) SetFormValidationErrors(err error) {
	for _, err := range err.(validator.ValidationErrors) {
		e.Errors = append(e.Errors, codes.ValidationError{FieldName: err.Field(), Message: err.Error()})
	}
}

func (e *ErrorResponse) SetValidationErrors(errors []*codes.ValidationError) {
	for _, err := range errors {
		e.Errors = append(e.Errors, codes.ValidationError{FieldName: err.Field(), Message: err.Error()})
	}
}

func NewBadRequest(err error) ErrorResponse {
	var errors []codes.ValidationError
	res := ErrorResponse{
		Code:    http.StatusBadRequest,
		Message: "Bad Request",
		Errors:  errors,
	}
	res.SetFormValidationErrors(err)
	return res
}

func NewValidationErrorBadRequest(validationErrors []*codes.ValidationError) ErrorResponse {
	var errors []codes.ValidationError
	res := ErrorResponse{
		Code:    http.StatusBadRequest,
		Message: "Bad Request",
		Errors:  errors,
	}
	res.SetValidationErrors(validationErrors)
	return res
}

func NewUnauthorized() ErrorResponse {
	var errors []codes.ValidationError
	res := ErrorResponse{
		Code:    http.StatusUnauthorized,
		Message: "unauthorized",
		Errors:  errors,
	}
	return res
}

func NewNotFound() ErrorResponse {
	var errors []codes.ValidationError
	res := ErrorResponse{
		Code:    http.StatusNotFound,
		Message: "Not Found",
		Errors:  errors,
	}
	return res
}

func NewConflic() ErrorResponse {
	var errors []codes.ValidationError
	res := ErrorResponse{
		Code:    http.StatusConflict,
		Message: "already exists",
		Errors:  errors,
	}
	return res
}

func NewInternalServerError() ErrorResponse {
	var errors []codes.ValidationError
	res := ErrorResponse{
		Code:    http.StatusInternalServerError,
		Message: "Internal Server Error",
		Errors:  errors,
	}
	return res
}
