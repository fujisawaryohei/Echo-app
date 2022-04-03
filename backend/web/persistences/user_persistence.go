package persistences

import (
	"errors"
	"fmt"

	"github.com/fujisawaryohei/blog-server/codes"
	"github.com/fujisawaryohei/blog-server/web/dto"
	"github.com/jackc/pgconn"
	"gorm.io/gorm"
)

type UserPersistence struct {
	dbConn *gorm.DB
}

func NewUserPersistence(db *gorm.DB) *UserPersistence {
	return &UserPersistence{
		dbConn: db,
	}
}

func (repo *UserPersistence) List() (*[]dto.User, error) {
	users := new([]dto.User)
	if err := repo.dbConn.Find(users).Error; err != nil {
		return users, fmt.Errorf("gateways/user.go List err: %w", err)
	}
	return users, nil
}

func (repo *UserPersistence) FindById(id int) (*dto.User, error) {
	user := new(dto.User)
	if err := repo.dbConn.First(user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, codes.ErrUserNotFound
		}
		return user, fmt.Errorf("gateways/user.go FindById err: %w", err)
	}
	return user, nil
}

func (repo *UserPersistence) FindByEmail(email string) (*dto.User, error) {
	user := new(dto.User)
	if err := repo.dbConn.First(user, "email=?", email).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, codes.ErrUserNotFound
		}
		return user, fmt.Errorf("gateway/user.go FindByEmail err: %w", err)
	}
	return user, nil
}

func (repo *UserPersistence) Save(userDTO *dto.User) error {
	if err := repo.dbConn.Create(userDTO).Error; err != nil {
		// https://github.com/go-gorm/gorm/issues/4135
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return codes.ErrUserEmailAlreadyExisted
		}
		return fmt.Errorf("gateway/user.go Save err: %w", err)
	}
	return nil
}

func (repo *UserPersistence) Update(id int, userDTO *dto.User) error {
	user, err := repo.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return codes.ErrUserNotFound
		}
		return fmt.Errorf("gateway/user.go Update err: %w", err)
	}

	if err := repo.dbConn.Model(user).Updates(userDTO).Error; err != nil {
		return fmt.Errorf("gateway/user.go Update err: %w", err)
	}
	return nil
}

func (repo *UserPersistence) Delete(id int) error {
	user := new(dto.User)
	if err := repo.dbConn.Delete(user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return codes.ErrUserNotFound
		}
		return fmt.Errorf("gateway/user.go Delete err: %w", err)
	}
	return nil
}
