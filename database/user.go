package database

import (
	"time"

	"github.com/fujisawaryohei/blog-server/web/dto"
	"github.com/fujisawaryohei/blog-server/web/utils"
)

type User struct {
	ID                   int `gorm:"primaryKey"`
	Name                 string
	Email                string
	Password             string
	PasswordConfirmation string
	CreatedAt            time.Time
	UpdatedAt            time.Time
}

func ConvertToUser(user *dto.User) *User {
	return &User{
		Name:                 user.Name,
		Email:                user.Email,
		Password:             utils.Hashed(user.Password),
		PasswordConfirmation: utils.Hashed(user.PasswordConfirmation),
	}
}
