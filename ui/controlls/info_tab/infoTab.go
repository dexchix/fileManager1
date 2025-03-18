package info_tab

import (
	"fileManager/ui"

	"github.com/gdamore/tcell/v2"
)

type InfoTab struct {
	State *InfoTabState
}

var styles ui.Styles = ui.GetStyles()

func (infoTab *InfoTab) RenderInfoTab(screen tcell.Screen) {
	screen.SetStyle(styles.DefaultStyle)
	for row := infoTab.State.Coordinates.Y1; row <= infoTab.State.Coordinates.Y2; row++ {
		for col := infoTab.State.Coordinates.X1; col <= infoTab.State.Coordinates.X2; col++ {
			screen.SetContent(col, row, ' ', nil, styles.StyleForView)
		}
	}
	infoTab.RenderBorder(screen)
	infoTab.FillInfoTab(screen)
}

func (infoTab *InfoTab) RenderBorder(screen tcell.Screen) {
	for col := infoTab.State.Coordinates.X1; col <= infoTab.State.Coordinates.X2; col++ {
		screen.SetContent(col, infoTab.State.Coordinates.Y1, tcell.RuneHLine, nil, styles.StyleForView)
		screen.SetContent(col, infoTab.State.Coordinates.Y2, tcell.RuneHLine, nil, styles.StyleForView)
	}

	for row := infoTab.State.Coordinates.Y1; row <= infoTab.State.Coordinates.Y2; row++ {
		screen.SetContent(infoTab.State.Coordinates.X1, row, tcell.RuneVLine, nil, styles.StyleForView)
		screen.SetContent(infoTab.State.Coordinates.X2, row, tcell.RuneVLine, nil, styles.StyleForView)
		if row == infoTab.State.Coordinates.Y1 {
			screen.SetContent(infoTab.State.Coordinates.X2/2+2, row, tcell.RuneTTee, nil, styles.StyleForView)
		} else if row == infoTab.State.Coordinates.Y2 {
			screen.SetContent(infoTab.State.Coordinates.X2/2+2, row, tcell.RuneBTee, nil, styles.StyleForView)
		} else {
			screen.SetContent(infoTab.State.Coordinates.X2/2+2, row, tcell.RuneVLine, nil, styles.StyleForView)
		}
	}

	screen.SetContent(infoTab.State.Coordinates.X1, infoTab.State.Coordinates.Y1, tcell.RuneULCorner, nil, styles.StyleForView)
	screen.SetContent(infoTab.State.Coordinates.X2, infoTab.State.Coordinates.Y1, tcell.RuneURCorner, nil, styles.StyleForView)
	screen.SetContent(infoTab.State.Coordinates.X1, infoTab.State.Coordinates.Y2, tcell.RuneLLCorner, nil, styles.StyleForView)
	screen.SetContent(infoTab.State.Coordinates.X2, infoTab.State.Coordinates.Y2, tcell.RuneLRCorner, nil, styles.StyleForView)
}

func (infoTab *InfoTab) FillInfoTab(screen tcell.Screen) {
	ui.DrawText(screen, infoTab.State.Y1da, infoTab.State.X1FC, infoTab.State.X2FC, infoTab.State.InfoCell, styles.StyleForFile)
	ui.DrawText(screen, infoTab.State.Y1da, infoTab.State.X1SC, infoTab.State.X2SC, infoTab.State.CommandCell, styles.StyleForFile)
}
