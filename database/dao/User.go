package dao

import (
	"time"

	"github.com/fujisawaryohei/echo-app/web/dto"
)

type User struct {
	ID        int `gorm:"primaryKey"`
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (User) ConvertToDAO(user *dto.User) *User {
	return &User{
		Name:  user.Name,
		Email: user.Email,
	}
}
