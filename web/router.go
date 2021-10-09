package web

import (
	"github.com/fujisawaryohei/blog-server/web/auth"
	"github.com/fujisawaryohei/blog-server/web/handlers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewServer(userHanlder *handlers.UserHandler) {
	e := echo.New()

	// アクセスロガー
	e.Use(middleware.Logger())

	config := middleware.JWTConfig{
		Claims:     &auth.JwtCustomClaim{},
		SigningKey: []byte("secret"),
	}

	// routing
	e.POST("/signup", userHanlder.Store)
	e.POST("/login", userHanlder.Login)

	r := e.Group("/users")
	r.Use(middleware.JWTWithConfig(config))
	r.GET("", userHanlder.List)
	r.GET("/:id", userHanlder.Find)
	r.PATCH("/:id", userHanlder.Update)
	r.DELETE("/:id", userHanlder.Delete)

	// サーバー起動
	e.Logger.Fatal(e.Start(":8080"))
}
