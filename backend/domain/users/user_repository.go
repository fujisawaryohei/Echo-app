package users

import (
	"github.com/fujisawaryohei/blog-server/web/dto"
)

type UserRepository interface {
	List() (*[]dto.User, error)
	FindById(id int) (*dto.User, error)
	FindByEmail(email string) (*dto.User, error)
	Save(userDTO *dto.User) error
	Update(id int, userDTO *dto.User) error
	Delete(id int) error
}
