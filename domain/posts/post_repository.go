package posts

import (
	"github.com/fujisawaryohei/blog-server/database"
	"github.com/fujisawaryohei/blog-server/web/dto"
)

type PostRepository interface {
	List() (*[]database.Post, error)
	FindById(id int) (*database.Post, error)
	Store(post *dto.Post) error
}
