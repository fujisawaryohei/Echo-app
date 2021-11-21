package dto

type Post struct {
	Title     string `json:"title" validate:"required,gte=0,lte=100"`
	Body      string `json:"body" validate:"required"`
	Published *bool  `json:"published" validate:"required"`
}
