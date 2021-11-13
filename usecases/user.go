package usecases

import (
	"errors"
	"fmt"
	"log"

	"github.com/fujisawaryohei/blog-server/codes"
	"github.com/fujisawaryohei/blog-server/database"
	"github.com/fujisawaryohei/blog-server/domain/users"
	"github.com/fujisawaryohei/blog-server/web/auth"
	"github.com/fujisawaryohei/blog-server/web/dto"
	"github.com/fujisawaryohei/blog-server/web/utils"
)

type UserUseCase struct {
	userRepository users.UserRepository
	authenticator  auth.IAuthenticator
}

func NewUserUsecase(repo users.UserRepository, authenticator auth.IAuthenticator) *UserUseCase {
	return &UserUseCase{
		userRepository: repo,
		authenticator:  authenticator,
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
func (u *UserUseCase) Store(userDTO *dto.User) (string, []*codes.ValidationError, error) {
	email := users.NewEmail(userDTO.Email)
	user, validationErrors := users.NewUser(userDTO.Name, email, userDTO.Password, userDTO.PasswordConfirmation)
	if len(validationErrors) != 0 {
		return "", validationErrors, nil
	}

	if err := u.userRepository.Save(user.ConvertToDTO()); err != nil {
		if errors.Is(err, codes.ErrUserEmailAlreadyExisted) {
			return "", nil, codes.ErrUserEmailAlreadyExisted
		}
		return "", nil, fmt.Errorf("usecases/user.go Store err: %w", err)
	}

	sigining_token, err := u.authenticator.GenerateToken(user.Email.Address)
	if err != nil {
		return "", nil, fmt.Errorf("usecases/user.go Store err: %w", err)
	}

	return sigining_token, nil, nil
}

func (u *UserUseCase) Login(loginUserDTO *dto.LoginUser) (string, error) {
	user, err := u.FindByEmail(loginUserDTO.Email)
	if err != nil {
		if errors.Is(err, codes.ErrUserNotFound) {
			return "", codes.ErrUserNotFound
		}
		log.Println(err.Error())
		return "", fmt.Errorf("usecases/user.go Login err: %w", err)
	}

	if err := utils.Compare(user.Password, loginUserDTO.Password); err != nil || user.Email != loginUserDTO.Email {
		return "", codes.ErrUserUnAuthorized
	}

	signing_token, err := u.authenticator.GenerateToken(loginUserDTO.Email)
	if err != nil {
		return "", fmt.Errorf("usecase/user.go Login err: %w", err)
	}

	return signing_token, nil
}

func (u *UserUseCase) Update(id int, userDTO *dto.User) ([]*codes.ValidationError, error) {
	email := users.NewEmail(userDTO.Email)
	user, validationErrors := users.NewUser(userDTO.Name, email, userDTO.Password, userDTO.PasswordConfirmation)
	if len(validationErrors) != 0 {
		return validationErrors, nil
	}

	if err := u.userRepository.Update(id, user.ConvertToDTO()); err != nil {
		if errors.Is(err, codes.ErrUserNotFound) {
			return nil, codes.ErrUserNotFound
		}
		return nil, fmt.Errorf("usecases/user.go Update err: %w", err)
	}
	return nil, nil
}

func (u *UserUseCase) Delete(id int) error {
	if err := u.userRepository.Delete(id); err != nil {
		if errors.Is(err, codes.ErrUserNotFound) {
			return codes.ErrUserNotFound
		}
		return fmt.Errorf("gateway/user.go Delete err: %w", err)
	}
	return nil
}
