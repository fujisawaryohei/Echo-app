package posts

type Post struct {
	Title     string
	Body      string
	Published bool
}

func NewPost(title string, body string, published bool) *Post {
	return &Post{
		Title:     title,
		Body:      body,
		Published: published,
	}
}
