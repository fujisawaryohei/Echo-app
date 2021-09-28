package usecases

import (
	"errors"
	"fmt"

	"github.com/fujisawaryohei/echo-app/codes"
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
		return users, fmt.Errorf("usecases/user.go list err: %s", err)
	}
	return users, err
}

func (u *UserUseCase) Find(id int) (*dao.User, error) {
	user, err := u.userRepository.FindById(id)
	if err != nil {
		if errors.Is(err, codes.ErrUserNotFound) {
			return user, codes.ErrUserNotFound
		}
		return user, fmt.Errorf("usecases/user.go Find err: %s", err)
	}
	return user, nil
}

func (u *UserUseCase) FindByEmail(email string) (*dao.User, error) {
	user, err := u.userRepository.FindByEmail(email)
	if err != nil {
		if errors.Is(err, codes.ErrUserNotFound) {
			return nil, codes.ErrUserNotFound
		}
		return nil, fmt.Errorf("usercases/user.go FindByEmail err: %s", err)
	}
	return user, err
}

func (u *UserUseCase) Store(user *dto.User) error {
	if err := u.userRepository.Save(user); err != nil {
		if errors.Is(err, codes.ErrUserAlreadyExisted) {
			return codes.ErrUserAlreadyExisted
		}
		return fmt.Errorf("usecases/user.go Store err: %s", err)
	}
	return nil
}

func (u *UserUseCase) Update(id int, newDTO *dto.User) error {
	if err := u.userRepository.Update(id, newDTO); err != nil {
		if errors.Is(err, codes.ErrUserNotFound) {
			return codes.ErrUserNotFound
		}
		return fmt.Errorf("usecases/user.go Update err: %s", err)
	}
	return nil
}

func (u *UserUseCase) Delete(id int) error {
	if err := u.userRepository.Delete(id); err != nil {
		return fmt.Errorf("gateway/user.go Delete err: %s", err)
	}
	return nil
}
