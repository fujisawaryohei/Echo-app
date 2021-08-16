package main

import (
	"github.com/fujisawaryohei/echo-app/web"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	web.NewServer(e)
	e.Logger.Fatal(e.Start(":8080"))
}
