package users

import (
	"github.com/fujisawaryohei/blog-server/codes"
	"github.com/fujisawaryohei/blog-server/web/dto"
)

type User struct {
	Name     string
	Email    Email
	Password Password
}

// TODO: refactor https://stackoverflow.com/questions/43336009/constructor-with-many-arguments
func NewUser(name string, email Email, password Password) (*User, []*codes.ValidationError) {
	user := &User{
		Name:     name,
		Email:    email,
		Password: password,
	}

	if !user.IsValid() {
		return nil, user.ValidationErrors()
	}

	return user, nil
}

func (u *User) IsValid() bool {
	if err := u.Password.PassowrdMatched(); err != nil {
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
	if err := u.Password.PassowrdMatched(); err != nil {
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

func (u *User) ConvertToDTO() *dto.User {
	return &dto.User{
		Name:                 u.Name,
		Email:                u.Email.Address,
		Password:             u.Password.Password,
		PasswordConfirmation: u.Password.PasswordConfirmation,
	}
}
