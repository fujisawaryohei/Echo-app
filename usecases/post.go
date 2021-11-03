package usecases

import (
	"fmt"

	"github.com/fujisawaryohei/blog-server/database"
	"github.com/fujisawaryohei/blog-server/domain/posts"
)

type PostUsecase struct {
	postRepository posts.PostRepository
}

func NewPostUsecase(repo posts.PostRepository) *PostUsecase {
	return &PostUsecase{
		postRepository: repo,
	}
}

func (u *PostUsecase) List() (*[]database.Post, error) {
	posts, err := u.postRepository.List()
	if err != nil {
		return posts, fmt.Errorf("usecase/post.go list err: %w", err)
	}
	return posts, err
}
