package ui

import "github.com/gdamore/tcell/v2"

type Styles struct {
	DefaultStyle         tcell.Style
	StyleForView         tcell.Style
	StyleForFile         tcell.Style
	StyleForDirectory    tcell.Style
	StyleForSelectedFile tcell.Style
}

func GetStyles() Styles {
	selectRowStyle := tcell.NewRGBColor(128, 128, 197)
	return Styles{
		DefaultStyle:         tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset),
		StyleForView:         tcell.StyleDefault.Background(tcell.ColorDarkBlue).Foreground(tcell.ColorLightGoldenrodYellow),
		StyleForFile:         tcell.StyleDefault.Background(tcell.ColorDarkBlue).Foreground(tcell.ColorWhite),
		StyleForDirectory:    tcell.StyleDefault.Background(tcell.ColorDarkBlue).Foreground(tcell.ColorWhite),
		StyleForSelectedFile: tcell.StyleDefault.Background(selectRowStyle).Foreground(tcell.ColorWhite),
	}
}
