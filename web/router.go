package web

import (
	"github.com/fujisawaryohei/blog-server/usecases"
	"github.com/fujisawaryohei/blog-server/web/auth"
	"github.com/fujisawaryohei/blog-server/web/handlers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewServer(userUseCase *usecases.UserUseCase) {
	e := echo.New()

	// アクセスロガー
	e.Use(middleware.Logger())

	config := middleware.JWTConfig{
		Claims:     &auth.JwtCustomClaim{},
		SigningKey: []byte("secret"),
	}

	// routing
	e.GET("/", handlers.Index)
	e.POST("/signup", handlers.StoreUser(userUseCase))
	e.POST("/login", handlers.Login(userUseCase))

	r := e.Group("/users")
	r.Use(middleware.JWTWithConfig(config))
	r.GET("", handlers.UserList(userUseCase))
	r.GET("/:id", handlers.FindUser(userUseCase))
	r.PATCH("/:id", handlers.UpdateUser(userUseCase))
	r.DELETE("/:id", handlers.DeleteUser(userUseCase))

	// サーバー起動
	e.Logger.Fatal(e.Start(":8080"))
}
