package posts

import "github.com/fujisawaryohei/blog-server/database"

type PostRepository interface {
	List() (*[]database.Post, error)
	FindById(id int) (*database.Post, error)
}
