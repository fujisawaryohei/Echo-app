package main

import (
	"fmt"

	"github.com/fujisawaryohei/blog-server/database"
	"github.com/fujisawaryohei/blog-server/database/seeds"
	"github.com/fujisawaryohei/blog-server/usecases"
	"github.com/fujisawaryohei/blog-server/web"
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
		userUseCase := usecases.NewUserUsecase(userRepository)
		userHandler := handlers.NewUserHandler(userUseCase)

		// サーバー起動
		web.NewServer(userHandler)
	}
}
