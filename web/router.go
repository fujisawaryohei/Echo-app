package web

import (
	"log"

	"github.com/fujisawaryohei/blog-server/web/auth"
	"github.com/fujisawaryohei/blog-server/web/handlers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewServer(userHanlder *handlers.UserHandler, postHandler *handlers.PostHandler) {
	e := echo.New()

	// アクセスロガー
	e.Use(middleware.Logger())

	signKey, err := auth.SignKey()
	if err != nil {
		log.Print(err)
	}

	config := middleware.JWTConfig{
		Claims:     &auth.JwtCustomClaim{},
		SigningKey: signKey,
	}

	// routing
	e.POST("/signup", userHanlder.Store)
	e.POST("/login", userHanlder.Login)

	users := e.Group("/users")
	users.Use(middleware.JWTWithConfig(config))
	users.GET("", userHanlder.List)
	users.GET("/:id", userHanlder.Find)
	users.PATCH("/:id", userHanlder.Update)
	users.DELETE("/:id", userHanlder.Delete)

	posts := e.Group("/posts")
	posts.Use(middleware.JWTWithConfig(config))
	posts.GET("", postHandler.List)
	posts.GET("/:id", postHandler.Find)
	posts.POST("", postHandler.Store)
	posts.PATCH("/:id", postHandler.Update)

	// サーバー起動
	e.Logger.Fatal(e.Start(":8080"))
}
