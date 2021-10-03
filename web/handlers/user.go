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
	"github.com/fujisawaryohei/echo-app/web/response"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

func UserList(usecase *usecases.UserUseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		users, err := usecase.List()
		if err != nil {
			log.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, response.NewInternalServerError())
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
				return c.JSON(http.StatusNotFound, response.NewNotFound())
			}
			log.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, response.NewInternalServerError())
		}
		return c.JSON(http.StatusOK, user)
	}
}

func StoreUser(usecase *usecases.UserUseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		userDTO := new(dto.User)
		if err := c.Bind(userDTO); err != nil {
			// TODO: いい感じにする
			return c.JSON(http.StatusBadRequest, "It contains an invalid value")
		}

		if err := validator.New().Struct(userDTO); err != nil {
			return c.JSON(http.StatusBadRequest, response.NewBadRequest(err))
		}

		if err := usecase.Store(userDTO); err != nil {
			if errors.Is(err, codes.ErrUserEmailAlreadyExisted) {
				return c.JSON(http.StatusConflict, response.NewConflic())
			}
			log.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, response.NewInternalServerError())
		}
		return c.JSON(http.StatusCreated, response.NewCreated())
	}
}

func UpdateUser(usecase *usecases.UserUseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		userDTO := new(dto.User)

		if err := c.Bind(userDTO); err != nil {
			// TODO: いい感じにする
			return c.JSON(http.StatusBadRequest, "It contains invalid Value")
		}

		if err := validator.New().Struct(userDTO); err != nil {
			return c.JSON(http.StatusBadRequest, response.NewBadRequest(err))
		}

		if err := usecase.Update(id, userDTO); err != nil {
			if errors.Is(err, codes.ErrUserNotFound) {
				return c.JSON(http.StatusNotFound, response.NewNotFound())
			}
			log.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, response.NewInternalServerError())
		}
		return c.JSON(http.StatusOK, response.NewSuccess())
	}
}

func DeleteUser(usecase *usecases.UserUseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		if err := usecase.Delete(id); err != nil {
			log.Println(err.Error())
			return c.JSON(http.StatusOK, response.NewInternalServerError())
		}
		return c.JSON(http.StatusOK, response.NewSuccess())
	}
}

func Login(usecase *usecases.UserUseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		loginUserDTO := new(dto.LoginUser)
		if err := c.Bind(loginUserDTO); err != nil {
			return c.JSON(http.StatusBadRequest, "It contains invalid Value")
		}

		if err := validator.New().Struct(loginUserDTO); err != nil {
			return c.JSON(http.StatusBadRequest, response.NewBadRequest(err))
		}

		if err := usecase.Login(loginUserDTO); err != nil {
			if errors.Is(err, codes.ErrUserNotFound) {
				return c.JSON(http.StatusNotFound, response.NewNotFound())
			}

			if errors.Is(err, codes.ErrUserUnAuthorized) {
				return c.JSON(http.StatusUnauthorized, response.NewUnauthorized())
			}
			log.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, response.NewInternalServerError())
		}

		signing_token, err := auth.GenerateToken(loginUserDTO.Email)
		if err != nil {
			return echo.ErrInternalServerError
		}

		return c.JSON(http.StatusOK, echo.Map{"access_token": signing_token})
	}
}
