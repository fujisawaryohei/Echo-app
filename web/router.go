package web

import (
	"github.com/fujisawaryohei/echo-app/web/handlers"
	"github.com/labstack/echo"
)

func NewServer(e *echo.Echo) {
	e.GET("/", handlers.Index)
}
