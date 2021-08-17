package gateways

import (
	"github.com/fujisawaryohei/echo-app/domain/repositories"
	"gorm.io/gorm"
)

func NewUserRepository(dbConn *gorm.DB) repositories.User {
	return &UserRepository{dbConn: dbConn}
}

type UserRepository struct {
	dbConn *gorm.DB
}

// func (repo *UserRepository) FindById(id string) entities.User {}

// func (repo *UserRepository) SaveUser() {}

// func (repo *UserRepository) Update() {}

// func (repo *UserRepository) Delete() {}
