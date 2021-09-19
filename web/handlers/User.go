package handlers

import (
	"net/http"
	"strconv"

	"github.com/fujisawaryohei/echo-app/usecases"
	"github.com/fujisawaryohei/echo-app/web/dto"
	"github.com/fujisawaryohei/echo-app/web/utils"
	"github.com/go-playground/validator/v10"
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
			// TODO: 何に対してのエラーハンドリングなのかを特定する
			return c.JSON(http.StatusBadRequest, "It contains an invalid value")
		}

		if err := validator.New().Struct(u); err != nil {
			errorRes := utils.NewBadRequestMessage(err)
			return c.JSON(errorRes.Code, errorRes)
		}

		if err := usecase.StoreUser(u); err != nil {
			errorRes := utils.NewInternalServerError(err)
			return c.JSON(errorRes.Code, errorRes)
		}

		return c.JSON(http.StatusCreated, utils.NewSuccessMessage())
	}
}

func DeleteUser(usecase *usecases.UserUseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		if err := usecase.DeleteUser(id); err != nil {
			return err
		}
		return c.JSON(http.StatusOK, utils.NewSuccessMessage())
	}
}
