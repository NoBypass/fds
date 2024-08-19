package version

import "github.com/labstack/gommon/color"

const VERSION = "v0.7.0"

var Banner = `:
   _______  ____  ____
  / __/ _ \/ __/ / __/__ _____  _____ ____
 / _// // /\ \  _\ \/ -_) __/ |/ / -_) __/
/_/ /____/___/ /___/\__/_/  |___/\__/_/   ` + color.Magenta(VERSION) + `
Backend API for all FDS services written in ` + color.CyanBg(color.White(" GO ")) + `
________________________________________________
`
