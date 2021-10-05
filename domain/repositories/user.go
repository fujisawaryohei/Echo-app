package repositories

import (
	"github.com/fujisawaryohei/blog-server/database"
	"github.com/fujisawaryohei/blog-server/web/dto"
)

type UserRepository interface {
	List() (*[]database.User, error)
	FindById(id int) (*database.User, error)
	FindByEmail(email string) (*database.User, error)
	Save(userDTO *dto.User) error
	Update(id int, userDTO *dto.User) error
	Delete(id int) error
}
