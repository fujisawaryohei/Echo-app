package handlers

import (
	"net/http"

	"github.com/fujisawaryohei/echo-app/usecases"
	"github.com/fujisawaryohei/echo-app/web/dto"
	"github.com/labstack/echo"
)

func StoreUser(usecase *usecases.UserUseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := new(dto.UserDTO)
		if err := c.Bind(u); err != nil {
			return err
		}
		usecase.StoreUser(u)
		return c.JSON(http.StatusOK, u)
	}
}
