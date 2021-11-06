package dto

type Post struct {
	Title     string `json:"title"`
	Body      string `json:"body"`
	Published bool   `json:"published"`
}
