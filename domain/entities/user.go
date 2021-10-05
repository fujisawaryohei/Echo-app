package entities

import (
	"github.com/fujisawaryohei/blog-server/web/dto"
)

type User struct {
	Name                 string
	Email                string
	Password             string
	PasswordConfirmation string
}

// TODO: refactor https://stackoverflow.com/questions/43336009/constructor-with-many-arguments
func NewUser(name string, email string, password string, passwordConfirmation string) *User {
	return &User{
		Name:                 name,
		Email:                email,
		Password:             password,
		PasswordConfirmation: passwordConfirmation,
	}
}

func (u *User) ConvertToDTO() *dto.User {
	return &dto.User{
		Name:                 u.Name,
		Email:                u.Email,
		Password:             u.Password,
		PasswordConfirmation: u.PasswordConfirmation,
	}
}
