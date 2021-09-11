package repositories

import (
	"github.com/fujisawaryohei/echo-app/database/dao"
	"github.com/fujisawaryohei/echo-app/web/dto"
)

type UserRepository interface {
	FindById(id string) (dao.User, error)
	SaveUser(user *dto.UserDTO) error
	Delete(id string) error
}
