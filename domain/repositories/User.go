package repositories

import "github.com/fujisawaryohei/echo-app/domain/entities"

type User interface {
	FindById(id string) entities.User
	SaveUser(*entities.User)
	Update(*entities.User)
	Delete(*entities.User)
}
