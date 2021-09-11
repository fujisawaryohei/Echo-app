package main

import (
	"fmt"

	"github.com/fujisawaryohei/echo-app/database"
	"github.com/fujisawaryohei/echo-app/database/seeds"
	"github.com/fujisawaryohei/echo-app/usecases"
	"github.com/fujisawaryohei/echo-app/web"
	"github.com/fujisawaryohei/echo-app/web/gateways"
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
		userRepository := gateways.NewUserRepository(db)
		userUseCase := usecases.NewUserUsecase(userRepository)

		// サーバー起動
		web.NewServer(userUseCase)
	}
}
