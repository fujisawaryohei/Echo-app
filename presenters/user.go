package presenters

import (
	"time"

	"github.com/fujisawaryohei/blog-server/database"
	viewmodel "github.com/fujisawaryohei/blog-server/presenters/viewModel"
)

func FormatUser(user *database.User) viewmodel.User {
	return viewmodel.User{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: createdAtFormatter(user.CreatedAt),
		UpdatedAt: updatedAtFormatter(user.UpdatedAt),
	}
}

func FormatUsers(users *[]database.User) []viewmodel.User {
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
