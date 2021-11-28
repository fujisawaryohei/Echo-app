package persistences

import (
	"errors"
	"fmt"

	"github.com/fujisawaryohei/blog-server/codes"
	"github.com/fujisawaryohei/blog-server/web/dto"
	"gorm.io/gorm"
)

type PostPersistence struct {
	dbConn *gorm.DB
}

func NewPostPersistence(db *gorm.DB) *PostPersistence {
	return &PostPersistence{
		dbConn: db,
	}
}

func (repo *PostPersistence) List() (*[]dto.Post, error) {
	posts := new([]dto.Post)
	if err := repo.dbConn.Find(posts).Error; err != nil {
		return posts, fmt.Errorf("gateways/post.go List err: %w", err)
	}
	return posts, nil
}

func (repo *PostPersistence) FindById(id int) (*dto.Post, error) {
	post := new(dto.Post)
	if err := repo.dbConn.First(post, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return post, codes.ErrPostNotFound
		}
		return post, fmt.Errorf("gateways/post.go FindById err: %w", err)
	}
	return post, nil
}

func (repo *PostPersistence) Store(postDTO *dto.Post) error {
	if err := repo.dbConn.Save(postDTO).Error; err != nil {
		return fmt.Errorf("gateway/post.go Save err: %w", err)
	}
	return nil
}

func (repo *PostPersistence) Update(id int, postDTO *dto.Post) error {
	post, err := repo.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return codes.ErrPostNotFound
		}
		return fmt.Errorf("gateay/post.go Update err: %w", err)
	}

	if err := repo.dbConn.Model(post).Updates(postDTO).Error; err != nil {
		return fmt.Errorf("gateway/post.go Update err: %w", err)
	}
	return nil
}

func (repo *PostPersistence) Delete(id int) error {
	post := new(dto.Post)
	if err := repo.dbConn.Delete(post, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return codes.ErrPostNotFound
		}
		return fmt.Errorf("gateway/post.go Delete err: %w", err)
	}
	return nil
}
