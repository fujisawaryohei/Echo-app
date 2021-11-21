package database

import (
	"time"

	"github.com/fujisawaryohei/blog-server/web/dto"
)

type Post struct {
	ID        int `gorm:"primaryKey"`
	Title     string
	Body      string `gorm:"type:text"`
	Published *bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func ConvertToPost(post *dto.Post) *Post {
	return &Post{
		Title:     post.Title,
		Body:      post.Body,
		Published: post.Published,
	}
}
