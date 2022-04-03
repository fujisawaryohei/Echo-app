package posts

import (
	"github.com/fujisawaryohei/blog-server/web/dto"
)

type PostRepository interface {
	List() (*[]dto.Post, error)
	FindById(id int) (*dto.Post, error)
	Store(post *dto.Post) error
	Update(id int, postDTO *dto.Post) error
	Delete(id int) error
}
