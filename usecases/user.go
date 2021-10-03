package usecases

import (
	"errors"
	"fmt"

	"github.com/fujisawaryohei/echo-app/codes"
	"github.com/fujisawaryohei/echo-app/database"
	"github.com/fujisawaryohei/echo-app/domain/entities"
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

func (u *UserUseCase) List() (*[]database.User, error) {
	users, err := u.userRepository.List()
	if err != nil {
		return users, fmt.Errorf("usecases/user.go list err: %w", err)
	}
	return users, err
}

func (u *UserUseCase) Find(id int) (*database.User, error) {
	user, err := u.userRepository.FindById(id)
	if err != nil {
		if errors.Is(err, codes.ErrUserNotFound) {
			return user, codes.ErrUserNotFound
		}
		return user, fmt.Errorf("usecases/user.go Find err: %w", err)
	}
	return user, nil
}

func (u *UserUseCase) FindByEmail(email string) (*database.User, error) {
	user, err := u.userRepository.FindByEmail(email)
	if err != nil {
		if errors.Is(err, codes.ErrUserNotFound) {
			return nil, codes.ErrUserNotFound
		}
		return nil, fmt.Errorf("usercases/user.go FindByEmail err: %w", err)
	}
	return user, err
}

// 原則レイヤ間のデータのやり取りはDTOを使用する。
// アプリケーション固有のロジックが発生した場合は、ドメインモデルを呼び出して処理してDTOに変換して別レイヤに渡す流れを取る。
func (u *UserUseCase) Store(userDTO *dto.User) error {
	user := entities.NewUser(userDTO.Name, userDTO.Email, userDTO.Password, userDTO.PasswordConfirmation)
	if err := u.userRepository.Save(user.ConvertToDTO()); err != nil {
		if errors.Is(err, codes.ErrUserEmailAlreadyExisted) {
			return codes.ErrUserEmailAlreadyExisted
		}
		return fmt.Errorf("usecases/user.go Store err: %w", err)
	}
	return nil
}

func (u *UserUseCase) Update(id int, userDTO *dto.User) error {
	user := entities.NewUser(userDTO.Name, userDTO.Email, userDTO.Password, userDTO.PasswordConfirmation)
	if err := u.userRepository.Update(id, user.ConvertToDTO()); err != nil {
		if errors.Is(err, codes.ErrUserNotFound) {
			return codes.ErrUserNotFound
		}
		return fmt.Errorf("usecases/user.go Update err: %w", err)
	}
	return nil
}

func (u *UserUseCase) Delete(id int) error {
	if err := u.userRepository.Delete(id); err != nil {
		return fmt.Errorf("gateway/user.go Delete err: %w", err)
	}
	return nil
}
