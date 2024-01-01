package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"server/internal/core/controller"
	"server/internal/core/middleware"
	"server/internal/pkg/conf"
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
________________________________________________
`)

	config := conf.ReadConfig()
	db := config.ConnectDB()

	dcc := controller.NewDiscordController(db)

	e.Use(middleware.Logger())
	e.Use(middleware.Timeout())

	discord := e.Group("/discord")
	discord.POST("/signup", dcc.Signup)
	discord.PATCH("/:id/daily", dcc.Daily)

	e.Logger.Fatal(e.Start(":8080"))
}
