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
	"github.com/labstack/echo"
)

type PostHandler struct {
	usecase *usecases.PostUsecase
}

func NewPostHanlder(usecase *usecases.PostUsecase) *PostHandler {
	return &PostHandler{
		usecase: usecase,
	}
}

func (h *PostHandler) List(c echo.Context) error {
	posts, err := h.usecase.List()
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, response.NewInternalServerError())
	}
	return c.JSON(http.StatusOK, posts)
}

func (h *PostHandler) Find(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	post, err := h.usecase.Find(id)
	if err != nil {
		if errors.Is(err, codes.ErrPostNotFound) {
			return c.JSON(http.StatusNotFound, response.NewNotFound())
		}
		return c.JSON(http.StatusInternalServerError, response.NewInternalServerError())
	}
	return c.JSON(http.StatusOK, post)
}

func (h *PostHandler) Store(c echo.Context) error {
	postDTO := new(dto.Post)
	if err := c.Bind(postDTO); err != nil {
		return c.JSON(http.StatusBadRequest, response.NewBadRequest)
	}

	if err := h.usecase.Store(postDTO); err != nil {
		return c.JSON(http.StatusInternalServerError, response.NewInternalServerError())
	}
	return c.JSON(http.StatusAccepted, nil)
}
