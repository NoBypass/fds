package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"server/internal/fds/middleware"
	"server/internal/fds/routes"
	"server/internal/pkg/consts"
)

const VERSION = "v0.2.0"

func main() {
	e := echo.New()

	e.HideBanner = true
	fmt.Println(`
   _______  ____  ____
  / __/ _ \/ __/ / __/__ _____  _____ ____
 / _// // /\ \  _\ \/ -_) __/ |/ / -_) __/
/_/ /____/___/ /___/\__/_/  |___/\__/_/   ` + consts.Purple.Sprint(VERSION) + `
Backend API for all FDS services written in ` + consts.WhiteOnCyan.Sprint(" GO ") + `
__________________________________________
`)

	e.Use(middleware.Logger())

	e.GET("/discord", routes.Discord)

	e.Logger.Fatal(e.Start(":8080"))
}
