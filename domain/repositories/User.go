package repositories

import (
	"github.com/fujisawaryohei/echo-app/database/dao"
	"github.com/fujisawaryohei/echo-app/web/dto"
)

type UserRepository interface {
	UserList() (*[]dao.User, error)
	FindById(id int) (*dao.User, error)
	SaveUser(user *dto.User) error
	Update(id int, newDTO *dto.User) error
	Delete(id int) error
}
