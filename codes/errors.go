package codes

import (
	"errors"
)

var (
	// 500系エラー
	ErrInternalServerError = errors.New("internal server error")

	// 400系エラー
	ErrUserNotFound            = errors.New("user not found")
	ErrUserUnAuthorized        = errors.New("user not unauthorized")
	ErrUserEmailAlreadyExisted = errors.New("email has already existed")
	ErrUserNameRequired        = errors.New("name is required")
	ErrUserEmailRequired       = errors.New("email is required")
	ErrUserNameTooLong         = errors.New("name is too long")
	ErrUserEmailTooLong        = errors.New("email is too long")
	ErrUserNameTooShort        = errors.New("name is too short")
	ErrUserEmailTooShort       = errors.New("email is too short")
	ErrPasswordRequired        = errors.New("password is required")
	ErrPasswordNotMatched      = errors.New("password_confirmation is not matched password")

	ErrPostNotFound = errors.New("post not found")
)

type ValidationError struct {
	FieldName string `json:"field"`
	Message   string `json:"message"`
}

func (v *ValidationError) Error() string {
	return v.Message
}

func (v *ValidationError) Field() string {
	return v.FieldName
}
