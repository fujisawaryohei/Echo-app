package routes

import (
	"github.com/fujisawaryohei/echo-app/web/controllers"
	"github.com/labstack/echo"
)

func Init(e *echo.Echo) {
	e.GET("/", controllers.Index)
}
