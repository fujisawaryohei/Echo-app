package gateways

import (
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

func (repo *UserRepository) UserList() (*[]dao.User, error) {
	usersDAO := new([]dao.User)
	if err := repo.dbConn.Find(usersDAO).Error; err != nil {
		return usersDAO, err
	}
	return usersDAO, nil
}

func (repo *UserRepository) FindById(id int) (*dao.User, error) {
	userDAO := new(dao.User)
	if err := repo.dbConn.First(userDAO, id).Error; err != nil {
		return userDAO, err
	}
	return userDAO, nil
}

func (repo *UserRepository) SaveUser(user *dto.User) error {
	userDAO := dao.User{}
	dao := userDAO.ConvertToDAO(user)
	if err := repo.dbConn.Create(dao).Error; err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) Update(id int, newDTO *dto.User) error {
	user, _ := repo.FindById(id)
	newDAO := dao.User{}.ConvertToDAO(newDTO)
	if err := repo.dbConn.Model(user).Updates(newDAO).Error; err != nil {
		return err
	}

	return nil
}

func (repo *UserRepository) Delete(id int) error {
	userDAO := new(dao.User)
	if err := repo.dbConn.Delete(userDAO, id).Error; err != nil {
		return err
	}
	return nil
}
