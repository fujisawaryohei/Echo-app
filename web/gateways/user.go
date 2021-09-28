package gateways

import (
	"errors"
	"fmt"

	"github.com/fujisawaryohei/echo-app/codes"
	"github.com/fujisawaryohei/echo-app/database/dao"
	"github.com/fujisawaryohei/echo-app/web/dto"
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

func (repo *UserRepository) List() (*[]dao.User, error) {
	usersDAO := new([]dao.User)
	if err := repo.dbConn.Find(usersDAO).Error; err != nil {
		return usersDAO, fmt.Errorf("gateways/user.go List err: %s", err)
	}
	return usersDAO, nil
}

func (repo *UserRepository) FindById(id int) (*dao.User, error) {
	userDAO := new(dao.User)
	if err := repo.dbConn.First(userDAO, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return userDAO, codes.ErrUserNotFound
		}
		return userDAO, fmt.Errorf("gateways/user.go FindById err: %s", err)
	}
	return userDAO, nil
}

func (repo *UserRepository) FindByEmail(email string) (*dao.User, error) {
	userDAO := new(dao.User)
	if err := repo.dbConn.First(userDAO, "email=?", email).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return userDAO, codes.ErrUserNotFound
		}
		return userDAO, fmt.Errorf("gateway/user.go FindByEmail err: %s", err)
	}
	return userDAO, nil
}

func (repo *UserRepository) Save(user *dto.User) error {
	userDAO := dao.User{}
	dao := userDAO.ConvertToDAO(user)
	if err := repo.dbConn.Create(dao).Error; err != nil {
		if errors.Is(err, gorm.ErrRegistered) {
			return codes.ErrUserAlreadyExisted
		}
		return fmt.Errorf("gateway/user.go Save err: %s", err)
	}
	return nil
}

func (repo *UserRepository) Update(id int, newDTO *dto.User) error {
	user, err := repo.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return codes.ErrUserNotFound
		}
		return fmt.Errorf("gateway/user.go Update err: %s", err)
	}

	newDAO := dao.User{}.ConvertToDAO(newDTO)
	if err := repo.dbConn.Model(user).Updates(newDAO).Error; err != nil {
		return fmt.Errorf("gateway/user.go Update err: %s", err)
	}
	return nil
}

func (repo *UserRepository) Delete(id int) error {
	userDAO := new(dao.User)
	if err := repo.dbConn.Delete(userDAO, id).Error; err != nil {
		return fmt.Errorf("gateway/user.go Delete err: %s", err)
	}
	return nil
}
