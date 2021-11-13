package users

import (
	"github.com/fujisawaryohei/blog-server/codes"
	"github.com/fujisawaryohei/blog-server/web/dto"
)

type UserFactory struct {
	UserRepository
}

func NewUserFactory(userRepository UserRepository) *UserFactory {
	return &UserFactory{
		UserRepository: userRepository,
	}
}

func (uf *UserFactory) Create(userDTO *dto.User) (*User, []*codes.ValidationError) {
	emailService := NewEmailService(uf.UserRepository)
	email := NewEmail(userDTO.Email, emailService)
	user, validationErrors := NewUser(userDTO.Name, email, userDTO.Password, userDTO.PasswordConfirmation)

	if len(validationErrors) != 0 {
		return nil, validationErrors
	}
	return user, nil
}
