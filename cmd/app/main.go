package main

import (
	"fmt"
	"github.com/NoBypass/fds/internal/app/controller"
	"github.com/NoBypass/fds/internal/app/middleware"
	"github.com/NoBypass/fds/internal/pkg/conf"
	"github.com/NoBypass/fds/internal/pkg/consts"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
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
	e.Use(middleware.Prepare(config))

	discord := e.Group("/discord")
	discord.Use(echojwt.WithConfig(*config.JWTConfig()))
	discord.POST("/signup", dcc.Signup)
	discord.PATCH("/:id/daily", dcc.Daily)
	discord.POST("/bot-login", dcc.BotLogin)

	e.Logger.Fatal(e.Start(":8080"))
}
