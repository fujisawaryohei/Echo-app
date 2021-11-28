package main

import (
	"fmt"

	"github.com/fujisawaryohei/blog-server/database"
	"github.com/fujisawaryohei/blog-server/database/seeds"
	"github.com/fujisawaryohei/blog-server/usecases"
	"github.com/fujisawaryohei/blog-server/web"
	"github.com/fujisawaryohei/blog-server/web/auth"
	"github.com/fujisawaryohei/blog-server/web/handlers"
	"github.com/fujisawaryohei/blog-server/web/persistences"
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
		userPersistence := persistences.NewUserPersistence(db)
		authenticator := auth.NewAuthenticator()
		userUseCase := usecases.NewUserUsecase(userPersistence, authenticator)
		userHandler := handlers.NewUserHandler(userUseCase)

		// PostHandler Init
		postPersistence := persistences.NewPostPersistence(db)
		postUseCase := usecases.NewPostUsecase(postPersistence)
		postHandler := handlers.NewPostHanlder(postUseCase)

		// サーバー起動
		web.NewServer(userHandler, postHandler)
	}
}
