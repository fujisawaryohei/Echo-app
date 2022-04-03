package presenters

import (
	"time"

	viewmodel "github.com/fujisawaryohei/blog-server/presenters/viewModel"
	"github.com/fujisawaryohei/blog-server/web/dto"
)

func CreateUserViewModel(user *dto.User) viewmodel.User {
	return viewmodel.User{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: createdAtFormatter(user.CreatedAt),
		UpdatedAt: updatedAtFormatter(user.UpdatedAt),
	}
}

func CreateUsersViewModel(users *[]dto.User) []viewmodel.User {
	var userViewModels []viewmodel.User
	for _, user := range *users {
		userViewModels = append(userViewModels, viewmodel.User{
			Id:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: createdAtFormatter(user.CreatedAt),
			UpdatedAt: updatedAtFormatter(user.UpdatedAt),
		})
	}
	return userViewModels
}

func createdAtFormatter(CreatedAt time.Time) string {
	return CreatedAt.Format("2006年01月02日 03時04分")
}

func updatedAtFormatter(UpdatedAt time.Time) string {
	return UpdatedAt.Format("2006年01月02日 03時04分")
}
