package users

import (
	"regexp"

	"github.com/fujisawaryohei/blog-server/codes"
)

type Email struct {
	Address string
}

func NewEmail(address string) *Email {
	email := &Email{Address: address}
	return email
}

func (e *Email) ValidFormat() *codes.ValidationError {
	pattern := `^(?i:[^ @"<>]+|".*")@(?i:[a-z1-9.])+.(?i:[a-z])+$`
	email_regexp := regexp.MustCompile(pattern)
	if len(email_regexp.FindAllString(e.Address, -1)) != 0 {
		return nil
	}
	return &codes.ValidationError{FieldName: "email", Message: codes.ErrUserEmailInvalidFormat.Error()}
}
