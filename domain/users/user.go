package users

import (
	"github.com/fujisawaryohei/blog-server/codes"
	"github.com/fujisawaryohei/blog-server/web/dto"
)

type User struct {
	Name                 string
	Email                Email
	Password             string
	PasswordConfirmation string
}

// TODO: refactor https://stackoverflow.com/questions/43336009/constructor-with-many-arguments
func NewUser(name string, email Email, password string, passwordConfirmation string) (*User, []*codes.ValidationError) {
	user := &User{
		Name:                 name,
		Email:                email,
		Password:             password,
		PasswordConfirmation: passwordConfirmation,
	}

	if !user.IsValid() {
		return nil, user.ValidationErrors()
	}

	return user, nil
}

func (u *User) IsValid() bool {
	if err := u.PassowrdMatched(); err != nil {
		return false
	}

	if err := u.Email.ValidFormat(); err != nil {
		return false
	}

	if err := u.Email.Duplicated(); err != nil {
		return false
	}
	return true
}

func (u *User) ValidationErrors() []*codes.ValidationError {
	var validationErrors []*codes.ValidationError
	if err := u.PassowrdMatched(); err != nil {
		validationErrors = append(validationErrors, err)
	}

	if err := u.Email.ValidFormat(); err != nil {
		validationErrors = append(validationErrors, err)
	}

	if err := u.Email.Duplicated(); err != nil {
		validationErrors = append(validationErrors, err)
	}
	return validationErrors
}

func (u *User) PassowrdMatched() *codes.ValidationError {
	if u.Password == u.PasswordConfirmation {
		return nil
	}
	return &codes.ValidationError{FieldName: "password", Message: codes.ErrPasswordNotMatched.Error()}
}

func (u *User) ConvertToDTO() *dto.User {
	return &dto.User{
		Name:                 u.Name,
		Email:                u.Email.Address,
		Password:             u.Password,
		PasswordConfirmation: u.PasswordConfirmation,
	}
}
