package usecases

import (
	"errors"
	"fmt"

	"github.com/fujisawaryohei/blog-server/codes"
	"github.com/fujisawaryohei/blog-server/database"
	"github.com/fujisawaryohei/blog-server/domain/posts"
	"github.com/fujisawaryohei/blog-server/web/dto"
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
		return posts, fmt.Errorf("usecases/post.go list err: %w", err)
	}
	return posts, err
}

func (u *PostUsecase) Find(id int) (*database.Post, error) {
	post, err := u.postRepository.FindById(id)
	if err != nil {
		if errors.Is(err, codes.ErrPostNotFound) {
			return post, codes.ErrPostNotFound
		}
		return post, fmt.Errorf("usecases/post.go Find err: %w", err)
	}
	return post, nil
}

func (u *PostUsecase) Store(postDTO *dto.Post) error {
	if err := u.postRepository.Store(postDTO); err != nil {
		return fmt.Errorf("usecases/post.go Store err: %w", err)
	}
	return nil
}

func (u *PostUsecase) Update(id int, postDTO *dto.Post) error {
	if err := u.postRepository.Update(id, postDTO); err != nil {
		return fmt.Errorf("usecases/post.go Store err: %w", err)
	}
	return nil
}
