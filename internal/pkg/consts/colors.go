package consts

import "github.com/fatih/color"

var (
	Purple        = color.New(color.FgMagenta)
	WhiteOnCyan   = color.New(color.BgHiCyan).Add(color.FgHiWhite)
	WhiteOnRed    = color.New(color.BgHiRed).Add(color.FgHiWhite)
	WhiteOnGreen  = color.New(color.BgHiGreen).Add(color.FgHiWhite)
	WhiteOnBlue   = color.New(color.BgHiBlue).Add(color.FgHiWhite)
	WhiteOnYellow = color.New(color.BgHiYellow).Add(color.FgHiWhite)
	Red           = color.New(color.FgHiRed)
	Green         = color.New(color.FgHiGreen)
	Yellow        = color.New(color.FgHiYellow)
)
