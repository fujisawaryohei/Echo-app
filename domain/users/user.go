package users

import (
	"github.com/fujisawaryohei/blog-server/codes"
	"github.com/fujisawaryohei/blog-server/web/dto"
)

type User struct {
	Name                 string
	Email                string
	Password             string
	PasswordConfirmation string
}

// TODO: refactor https://stackoverflow.com/questions/43336009/constructor-with-many-arguments
func NewUser(name string, email string, password string, passwordConfirmation string) (*User, []error) {
	user := &User{
		Name:                 name,
		Email:                email,
		Password:             password,
		PasswordConfirmation: passwordConfirmation,
	}

	if user.IsValid() {
		return user, nil
	}

	return nil, user.Errors()
}

func (u *User) ConvertToDTO() *dto.User {
	return &dto.User{
		Name:                 u.Name,
		Email:                u.Email,
		Password:             u.Password,
		PasswordConfirmation: u.PasswordConfirmation,
	}
}

func (u *User) IsValid() bool {
	var errors []error
	if u.Password == u.PasswordConfirmation {
		errors = append(errors, codes.ErrPasswordNotMatched)
	}

	if len(errors) == 0 {
		return false
	}

	return true
}

func (u *User) Errors() []error {
	var errors []error

	if u.Password == u.PasswordConfirmation {
		errors = append(errors, codes.ErrPasswordNotMatched)
	}
	return errors
}
