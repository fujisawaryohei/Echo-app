package gateways

import (
	"errors"
	"fmt"

	"github.com/fujisawaryohei/echo-app/codes"
	"github.com/fujisawaryohei/echo-app/database"
	"github.com/fujisawaryohei/echo-app/web/dto"
	"github.com/jackc/pgconn"
	"gorm.io/gorm"
)

type UserRepository struct {
	dbConn *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		dbConn: db,
	}
}

func (repo *UserRepository) List() (*[]database.User, error) {
	users := new([]database.User)
	if err := repo.dbConn.Find(users).Error; err != nil {
		return users, fmt.Errorf("gateways/user.go List err: %w", err)
	}
	return users, nil
}

func (repo *UserRepository) FindById(id int) (*database.User, error) {
	user := new(database.User)
	if err := repo.dbConn.First(user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, codes.ErrUserNotFound
		}
		return user, fmt.Errorf("gateways/user.go FindById err: %w", err)
	}
	return user, nil
}

func (repo *UserRepository) FindByEmail(email string) (*database.User, error) {
	user := new(database.User)
	if err := repo.dbConn.First(user, "email=?", email).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, codes.ErrUserNotFound
		}
		return user, fmt.Errorf("gateway/user.go FindByEmail err: %w", err)
	}
	return user, nil
}

func (repo *UserRepository) Save(userDTO *dto.User) error {
	user := database.ConvertToUser(userDTO)
	if err := repo.dbConn.Create(user).Error; err != nil {
		// TODO: errors.As について調べてリファクタする
		// https://github.com/go-gorm/gorm/issues/4135
		var pgErr *pgconn.PgError
		errors.As(err, &pgErr)
		if pgErr.Code == "23505" {
			return codes.ErrUserEmailAlreadyExisted
		}
		fmt.Println(pgErr.Code)
		return fmt.Errorf("gateway/user.go Save err: %w", err)
	}
	return nil
}

func (repo *UserRepository) Update(id int, userDTO *dto.User) error {
	user, err := repo.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return codes.ErrUserNotFound
		}
		return fmt.Errorf("gateway/user.go Update err: %w", err)
	}

	newUser := database.ConvertToUser(userDTO)
	if err := repo.dbConn.Model(user).Updates(newUser).Error; err != nil {
		return fmt.Errorf("gateway/user.go Update err: %w", err)
	}
	return nil
}

func (repo *UserRepository) Delete(id int) error {
	user := new(database.User)
	if err := repo.dbConn.Delete(user, id).Error; err != nil {
		return fmt.Errorf("gateway/user.go Delete err: %w", err)
	}
	return nil
}
