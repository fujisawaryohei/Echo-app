package usecases

import (
	"github.com/fujisawaryohei/echo-app/domain/repositories"
	"github.com/fujisawaryohei/echo-app/web/dto"
)

type UserUseCase struct {
	userRepository repositories.UserRepository
}

func NewUserUsecase(repo repositories.UserRepository) *UserUseCase {
	return &UserUseCase{
		userRepository: repo,
	}
}

func (u *UserUseCase) StoreUser(user *dto.UserDTO) error {
	if err := u.userRepository.SaveUser(user); err != nil {
		return err
	}
	return nil
}
