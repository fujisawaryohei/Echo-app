package handlers

import (
	"log"
	"net/http"

	"github.com/fujisawaryohei/blog-server/usecases"
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
