package dao

import (
	"time"

	"github.com/fujisawaryohei/echo-app/web/dto"
)

type User struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Email      string
	Created_at time.Time
	Updated_at time.Time
}

func (User) ConvertToDAO(user *dto.UserDTO) *User {
	return &User{
		Name:  user.Name,
		Email: user.Email,
	}
}
