package web

import (
	"github.com/fujisawaryohei/echo-app/usecases"
	"github.com/fujisawaryohei/echo-app/web/handlers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewServer(userUseCase *usecases.UserUseCase) {
	e := echo.New()

	// アクセスロガー
	e.Use(middleware.Logger())

	// routing
	e.GET("/", handlers.Index)
	e.GET("/users/:id", handlers.Find(userUseCase))
	e.POST("/users", handlers.StoreUser(userUseCase))

	// サーバー起動
	e.Logger.Fatal(e.Start(":8080"))
}
