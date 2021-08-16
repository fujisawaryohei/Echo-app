package main

import (
	"github.com/fujisawaryohei/echo-app/web/routes"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	routes.Init(e)
	e.Logger.Fatal(e.Start(":8080"))
}
