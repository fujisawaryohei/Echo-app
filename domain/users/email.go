package users

import (
	"regexp"

	"github.com/fujisawaryohei/blog-server/codes"
)

type Email struct {
	Address string
	Service EmailService
}

func NewEmail(address string, service EmailService) *Email {
	email := &Email{
		Address: address,
		Service: service,
	}
	return email
}

func (e *Email) ValidFormat() *codes.ValidationError {
	ADDRESS_PATTERN := `^(?i:[^ @"<>]+|".*")@(?i:[a-z1-9.])+.(?i:[a-z])+$`
	email_regexp := regexp.MustCompile(ADDRESS_PATTERN)
	if len(email_regexp.FindAllString(e.Address, -1)) != 0 {
		return nil
	}
	return &codes.ValidationError{FieldName: "email", Message: codes.ErrUserEmailInvalidFormat.Error()}
}

func (e *Email) Duplicated() *codes.ValidationError {
	if !e.Service.Duplicated(e.Address) {
		return nil
	}
	return &codes.ValidationError{FieldName: "email", Message: codes.ErrUserEmailAlreadyExisted.Error()}
}
