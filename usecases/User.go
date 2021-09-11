package usecases

import (
	"github.com/fujisawaryohei/echo-app/database/dao"
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

func (u *UserUseCase) Find(id int) (*dao.User, error) {
	user, err := u.userRepository.FindById(id)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *UserUseCase) StoreUser(user *dto.User) error {
	if err := u.userRepository.SaveUser(user); err != nil {
		return err
	}
	return nil
}
