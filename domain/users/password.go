package users

import "github.com/fujisawaryohei/blog-server/codes"

type Password struct {
	Password             string
	PasswordConfirmation string
}

func NewPassword(password string, passwordConfirmation string) Password {
	return Password{
		Password:             password,
		PasswordConfirmation: passwordConfirmation,
	}
}

func (p *Password) PassowrdMatched() *codes.ValidationError {
	if p.Password == p.PasswordConfirmation {
		return nil
	}
	return &codes.ValidationError{FieldName: "password", Message: codes.ErrPasswordNotMatched.Error()}
}
