package repositories

import (
	"github.com/fujisawaryohei/echo-app/database/dao"
	"github.com/fujisawaryohei/echo-app/web/dto"
)

type UserRepository interface {
	FindById(id int) (*dao.User, error)
	SaveUser(user *dto.User) error
	Delete(id int) error
}
