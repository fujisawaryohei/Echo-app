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
func NewUser(name string, email string, password string, passwordConfirmation string) (*User, []*codes.ValidationError) {
	user := &User{
		Name:                 name,
		Email:                email,
		Password:             password,
		PasswordConfirmation: passwordConfirmation,
	}

	ValidationErrors := user.IsValid()
	if len(ValidationErrors) != 0 {
		return nil, ValidationErrors
	}

	return user, nil
}

func (u *User) ConvertToDTO() *dto.User {
	return &dto.User{
		Name:                 u.Name,
		Email:                u.Email,
		Password:             u.Password,
		PasswordConfirmation: u.PasswordConfirmation,
	}
}

func (u *User) IsValid() []*codes.ValidationError {
	var validationErrors []*codes.ValidationError
	if err := u.PassowrdMatched(); err != nil {
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
