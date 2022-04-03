package users

import (
	"errors"

	"github.com/fujisawaryohei/blog-server/codes"
)

type EmailService struct {
	UserRepository
}

func NewEmailService(userRepository UserRepository) EmailService {
	return EmailService{
		UserRepository: userRepository,
	}
}

func (e *EmailService) Duplicated(address string) bool {
	if _, err := e.UserRepository.FindByEmail(address); err != nil {
		if errors.Is(err, codes.ErrUserNotFound) {
			return false
		}
	}
	return true
}
