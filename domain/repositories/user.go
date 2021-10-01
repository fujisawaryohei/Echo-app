package repositories

import (
	"github.com/fujisawaryohei/echo-app/database"
	"github.com/fujisawaryohei/echo-app/web/dto"
)

type UserRepository interface {
	List() (*[]database.User, error)
	FindById(id int) (*database.User, error)
	FindByEmail(email string) (*database.User, error)
	Save(userDTO *dto.User) error
	Update(id int, userDTO *dto.User) error
	Delete(id int) error
}
