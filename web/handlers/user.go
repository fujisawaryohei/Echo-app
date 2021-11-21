package handlers

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/fujisawaryohei/blog-server/codes"
	"github.com/fujisawaryohei/blog-server/usecases"
	"github.com/fujisawaryohei/blog-server/web/dto"
	"github.com/fujisawaryohei/blog-server/web/response"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

type UserHandler struct {
	usecase *usecases.UserUseCase
}

func NewUserHandler(usecase *usecases.UserUseCase) *UserHandler {
	return &UserHandler{
		usecase: usecase,
	}
}

func (h *UserHandler) List(c echo.Context) error {
	users, err := h.usecase.List()
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, response.NewInternalServerError())
	}
	return c.JSON(http.StatusOK, users)
}

func (h *UserHandler) Find(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := h.usecase.Find(id)
	if err != nil {
		if errors.Is(err, codes.ErrUserNotFound) {
			return c.JSON(http.StatusNotFound, response.NewNotFound())
		}
		log.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, response.NewInternalServerError())
	}
	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) Store(c echo.Context) error {
	userDTO := new(dto.User)
	if err := c.Bind(userDTO); err != nil {
		// TODO: いい感じにする
		return c.JSON(http.StatusBadRequest, "It contains an invalid value")
	}

	if err := validator.New().Struct(userDTO); err != nil {
		return c.JSON(http.StatusBadRequest, response.NewBadRequest(err))
	}

	signing_token, validationErrors, err := h.usecase.Store(userDTO)
	if len(validationErrors) != 0 {
		return c.JSON(http.StatusBadRequest, response.NewValidationErrorBadRequest(validationErrors))
	}

	if err != nil {
		if errors.Is(err, codes.ErrUserEmailAlreadyExisted) {
			return c.JSON(http.StatusConflict, response.NewConflic())
		}
		log.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, response.NewInternalServerError())
	}
	return c.JSON(http.StatusOK, echo.Map{"access_token": signing_token})
}

func (h *UserHandler) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	userDTO := new(dto.User)
	if err := c.Bind(userDTO); err != nil {
		// TODO: いい感じにする
		return c.JSON(http.StatusBadRequest, "It contains invalid Value")
	}

	if err := validator.New().Struct(userDTO); err != nil {
		return c.JSON(http.StatusBadRequest, response.NewBadRequest(err))
	}

	validationErrors, err := h.usecase.Update(id, userDTO)
	if len(validationErrors) != 0 {
		return c.JSON(http.StatusBadRequest, response.NewValidationErrorBadRequest(validationErrors))
	}

	if err != nil {
		if errors.Is(err, codes.ErrUserNotFound) {
			return c.JSON(http.StatusNotFound, response.NewNotFound())
		}
		log.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, response.NewInternalServerError())
	}
	return c.JSON(http.StatusAccepted, nil)
}

func (h *UserHandler) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.usecase.Delete(id); err != nil {
		if errors.Is(err, codes.ErrUserNotFound) {
			return c.JSON(http.StatusNotFound, response.NewNotFound())
		}
		log.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, response.NewInternalServerError())
	}
	return c.JSON(http.StatusAccepted, nil)
}

func (h *UserHandler) Login(c echo.Context) error {
	loginUserDTO := new(dto.LoginUser)
	if err := c.Bind(loginUserDTO); err != nil {
		return c.JSON(http.StatusBadRequest, "It contains invalid Value")
	}

	if err := validator.New().Struct(loginUserDTO); err != nil {
		return c.JSON(http.StatusBadRequest, response.NewBadRequest(err))
	}

	signing_token, err := h.usecase.Login(loginUserDTO)
	if err != nil {
		if errors.Is(err, codes.ErrUserNotFound) {
			return c.JSON(http.StatusNotFound, response.NewNotFound())
		}
		if errors.Is(err, codes.ErrUserUnAuthorized) {
			return c.JSON(http.StatusUnauthorized, response.NewUnauthorized())
		}
		log.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, response.NewInternalServerError())
	}

	return c.JSON(http.StatusOK, echo.Map{"access_token": signing_token})
}
