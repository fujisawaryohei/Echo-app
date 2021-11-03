package main

import (
	"fmt"

	"github.com/fujisawaryohei/blog-server/database"
	"github.com/fujisawaryohei/blog-server/database/seeds"
	"github.com/fujisawaryohei/blog-server/usecases"
	"github.com/fujisawaryohei/blog-server/web"
	"github.com/fujisawaryohei/blog-server/web/auth"
	"github.com/fujisawaryohei/blog-server/web/gateways"
	"github.com/fujisawaryohei/blog-server/web/handlers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	// アクセスロガー
	e.Use(middleware.Logger())

	// DB Init
	db, err := database.NewConnection()
	if err != nil {
		fmt.Printf("got err, %+v\n", err)
	}

	// Seed Command Handler
	commandArgs := seeds.SeedHandler(db)

	// App Init
	if len(commandArgs) == 0 {
		// UserHandler Init
		userRepository := gateways.NewUserRepository(db)
		authenticator := auth.NewAuthenticator()
		userUseCase := usecases.NewUserUsecase(userRepository, authenticator)
		userHandler := handlers.NewUserHandler(userUseCase)

		// PostHandler Init
		postRepository := gateways.NewPostRepository(db)
		postUseCase := usecases.NewPostUsecase(postRepository)
		postHandler := handlers.NewPostHanlder(postUseCase)

		// サーバー起動
		web.NewServer(userHandler, postHandler)
	}
}
