package main

import (
	"github.com/NoBypass/fds/internal/backend/startup"
	"github.com/NoBypass/fds/internal/frontend/renderer"
	"github.com/NoBypass/fds/internal/pkg/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/color"
	"github.com/labstack/gommon/log"
)

func main() {
	e := echo.New()
	e.HideBanner = true
	e.Logger.SetLevel(log.INFO)

	e.Logger.Print(`:
   _______  ____  ____
  / __/ _ \/ __/ / __/__ _____  _____ ____
 / _// // /\ \  _\ \/ -_) __/ |/ / -_) __/
/_/ /____/___/ /___/\__/_/  |___/\__/_/   ` + color.Magenta(utils.VERSION) + `
Backend API for all FDS services written in ` + color.CyanBg(color.White(" GO ")) + `
________________________________________________
`)

	app := startup.Connect(e)
	defer app.Closer.Close()
	e.Debug = app.Cfg.Development != ""

	static := renderer.New()
	e.Renderer = static

	startup.Router(app)

	e.Logger.Fatal(e.Start(":8080"))
}
