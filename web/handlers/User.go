package handlers

import (
	"net/http"
	"strconv"

	"github.com/fujisawaryohei/echo-app/usecases"
	"github.com/fujisawaryohei/echo-app/web/dto"
	"github.com/labstack/echo"
)

type SuccessMsg struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

func Find(usecase *usecases.UserUseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		user, err := usecase.Find(id)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, user)
	}
}

func StoreUser(usecase *usecases.UserUseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := new(dto.User)
		if err := c.Bind(u); err != nil {
			return err
		}
		usecase.StoreUser(u)
		return c.JSON(http.StatusOK, u)
	}
}

func DeleteUser(usecase *usecases.UserUseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		if err := usecase.DeleteUser(id); err != nil {
			return err
		}
		SuccessMsg := &SuccessMsg{StatusCode: http.StatusOK, Message: "削除が完了しました"}
		return c.JSON(http.StatusOK, SuccessMsg)
	}
}
