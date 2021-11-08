package gateways

import (
	"errors"
	"fmt"

	"github.com/fujisawaryohei/blog-server/codes"
	"github.com/fujisawaryohei/blog-server/database"
	"github.com/fujisawaryohei/blog-server/web/dto"
	"gorm.io/gorm"
)

// TODO: ドメインモデルに変換して返す
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

func (repo *PostRepository) FindById(id int) (*database.Post, error) {
	post := new(database.Post)
	if err := repo.dbConn.First(post, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return post, codes.ErrPostNotFound
		}
		return post, fmt.Errorf("gateways/post.go FindById err: %w", err)
	}
	return post, nil
}

func (repo *PostRepository) Store(postDTO *dto.Post) error {
	post := database.ConvertToPost(postDTO)
	if err := repo.dbConn.Save(post).Error; err != nil {
		return fmt.Errorf("gateway/post.go Save err: %w", err)
	}
	return nil
}
