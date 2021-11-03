package gateways

import (
	"fmt"

	"github.com/fujisawaryohei/blog-server/database"
	"gorm.io/gorm"
)

type PostRepository struct {
	dbConn *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{
		dbConn: db,
	}
}

func (repo *PostRepository) List() (*[]database.Post, error) {
	posts := new([]database.Post)
	if err := repo.dbConn.Find(posts).Error; err != nil {
		return posts, fmt.Errorf("gateways/post.go List err: %w", err)
	}
	return posts, nil
}
