package dto

import "time"

type Post struct {
	ID        int       `json:"id"`
	Title     string    `json:"title" validate:"required,gte=0,lte=100"`
	Body      string    `json:"body" validate:"required"`
	Published *bool     `json:"published" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
