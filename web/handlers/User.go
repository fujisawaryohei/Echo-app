package handlers

import (
	"net/http"
	"strconv"

	"github.com/fujisawaryohei/echo-app/usecases"
	"github.com/fujisawaryohei/echo-app/web/auth"
	"github.com/fujisawaryohei/echo-app/web/dto"
	"github.com/fujisawaryohei/echo-app/web/utils"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

func UserList(usecase *usecases.UserUseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		// TODO: エラーハンドリングの対象としての想定は401なので一旦必要なし
		users, _ := usecase.List()
		return c.JSON(http.StatusOK, users)
	}
}

func FindUser(usecase *usecases.UserUseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		user, err := usecase.Find(id)
		if err != nil {
			errorRes := utils.NewNotFoundMessage(err)
			return c.JSON(errorRes.Code, errorRes)
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

		if err := usecase.Store(u); err != nil {
			errorRes := utils.NewInternalServerError(err)
			return c.JSON(errorRes.Code, errorRes)
		}

		return c.JSON(http.StatusCreated, utils.NewSuccessMessage())
	}
}

func UpdateUser(usecase *usecases.UserUseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		newDTO := new(dto.User)

		if err := c.Bind(newDTO); err != nil {
			// TODO: 何に対してのエラーハンドリングなのかを特定する
			return c.JSON(http.StatusBadRequest, "It contains invalid Value")
		}

		if err := usecase.Update(id, newDTO); err != nil {
			errorRes := utils.NewNotFoundMessage(err)
			return c.JSON(errorRes.Code, errorRes)
		}

		return c.JSON(http.StatusOK, utils.NewSuccessMessage())
	}
}

func DeleteUser(usecase *usecases.UserUseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		if err := usecase.Delete(id); err != nil {
			errorRes := utils.NewInternalServerError(err)
			return c.JSON(errorRes.Code, errorRes)
		}

		return c.JSON(http.StatusOK, utils.NewSuccessMessage())
	}
}

func Login(usecase *usecases.UserUseCase) echo.HandlerFunc {
	u := new(dto.LoginUser)
	return func(c echo.Context) error {
		if err := c.Bind(u); err != nil {
			return c.JSON(http.StatusBadRequest, "It contains invalid Value")
		}

		userDAO, err := usecase.FindByEmail(u.Email)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.ErrNotFound)
		}

		if err := utils.Compare(userDAO.Password, u.Password); err != nil || userDAO.Email != u.Email {
			return c.JSON(http.StatusUnauthorized, echo.ErrUnauthorized)
		}

		signing_token, err := auth.GenerateToken(u.Email)
		if err != nil {
			return echo.ErrInternalServerError
		}

		return c.JSON(http.StatusOK, echo.Map{"token": signing_token})
	}
}
