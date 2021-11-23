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

	admin := e.Group("/admin")
	admin.Use(middleware.JWTWithConfig(config))
	admin.GET("/users", userHanlder.List)
	admin.GET("/users/:id", userHanlder.Find)
	admin.PATCH("/users/:id", userHanlder.Update)
	admin.DELETE("/users/:id", userHanlder.Delete)
	admin.POST("/posts", postHandler.Store)
	admin.PATCH("/posts/:id", postHandler.Update)
	admin.DELETE("/posts/:id", postHandler.Delete)
	e.GET("/posts", postHandler.List)
	e.GET("/posts/:id", postHandler.Find)

	// サーバー起動
	e.Logger.Fatal(e.Start(":8080"))
}
