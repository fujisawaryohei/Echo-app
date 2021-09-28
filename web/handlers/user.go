package handlers

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/fujisawaryohei/echo-app/codes"
	"github.com/fujisawaryohei/echo-app/usecases"
	"github.com/fujisawaryohei/echo-app/web/auth"
	"github.com/fujisawaryohei/echo-app/web/dto"
	"github.com/fujisawaryohei/echo-app/web/utils"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

func UserList(usecase *usecases.UserUseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		users, err := usecase.List()
		if err != nil {
			log.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, utils.NewInternalServerError())
		}
		return c.JSON(http.StatusOK, users)
	}
}

func FindUser(usecase *usecases.UserUseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		user, err := usecase.Find(id)
		if err != nil {
			if errors.Is(err, codes.ErrUserNotFound) {
				return c.JSON(http.StatusNotFound, utils.NewNotFoundMessage())
			}
			log.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, utils.NewInternalServerError())
		}
		return c.JSON(http.StatusOK, user)
	}
}

func StoreUser(usecase *usecases.UserUseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := new(dto.User)
		if err := c.Bind(u); err != nil {
			// TODO: いい感じにする
			return c.JSON(http.StatusBadRequest, "It contains an invalid value")
		}

		if err := validator.New().Struct(u); err != nil {
			return c.JSON(http.StatusBadRequest, utils.NewBadRequestMessage(err))
		}

		if err := usecase.Store(u); err != nil {
			if errors.Is(err, codes.ErrUserAlreadyExisted) {
				return c.JSON(http.StatusConflict, utils.NewConflic())
			}
			log.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, utils.NewInternalServerError())
		}
		return c.JSON(http.StatusCreated, utils.NewSuccessMessage())
	}
}

func UpdateUser(usecase *usecases.UserUseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		newDTO := new(dto.User)

		if err := c.Bind(newDTO); err != nil {
			// TODO: いい感じにする
			return c.JSON(http.StatusBadRequest, "It contains invalid Value")
		}

		if err := usecase.Update(id, newDTO); err != nil {
			if errors.Is(err, codes.ErrUserNotFound) {
				return c.JSON(http.StatusNotFound, utils.NewNotFoundMessage())
			}
			log.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, utils.NewInternalServerError())
		}
		return c.JSON(http.StatusOK, utils.NewSuccessMessage())
	}
}

func DeleteUser(usecase *usecases.UserUseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		if err := usecase.Delete(id); err != nil {
			log.Println(err.Error())
			return c.JSON(http.StatusOK, utils.NewInternalServerError())
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
			if errors.Is(err, codes.ErrUserNotFound) {
				return c.JSON(http.StatusBadRequest, utils.NewNotFoundMessage())
			}
			log.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, utils.NewInternalServerError())
		}

		if err := utils.Compare(userDAO.Password, u.Password); err != nil || userDAO.Email != u.Email {
			return c.JSON(http.StatusUnauthorized, utils.NewUnauthorized())
		}

		signing_token, err := auth.GenerateToken(u.Email)
		if err != nil {
			return echo.ErrInternalServerError
		}

		return c.JSON(http.StatusOK, echo.Map{"token": signing_token})
	}
}
