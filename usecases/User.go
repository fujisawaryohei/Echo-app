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

func (u *UserUseCase) List() (*[]dao.User, error) {
	users, err := u.userRepository.List()
	if err != nil {
		return users, err
	}
	return users, err
}

func (u *UserUseCase) Find(id int) (*dao.User, error) {
	user, err := u.userRepository.FindById(id)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *UserUseCase) Store(user *dto.User) error {
	if err := u.userRepository.Save(user); err != nil {
		return err
	}
	return nil
}

func (u *UserUseCase) Update(id int, newDTO *dto.User) error {
	if err := u.userRepository.Update(id, newDTO); err != nil {
		return err
	}
	return nil
}

func (u *UserUseCase) Delete(id int) error {
	if err := u.userRepository.Delete(id); err != nil {
		return err
	}
	return nil
}
