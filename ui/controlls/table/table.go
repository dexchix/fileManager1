package table

import (
	"fileManager/models"
	"fileManager/ui"

	"github.com/gdamore/tcell/v2"
)

type Table struct {
	State *TableState
}

type Row struct {
	Y    int
	File models.File
}

var styles ui.Styles = ui.GetStyles()

func (table *Table) RenderTable(screen tcell.Screen) {
	screen.SetStyle(styles.DefaultStyle)
	for row := table.State.Coordinates.Y1; row <= table.State.Coordinates.Y2; row++ {
		for col := table.State.Coordinates.X1; col <= table.State.Coordinates.X2; col++ {
			screen.SetContent(col, row, ' ', nil, styles.StyleForView)
		}
	}
	table.RenderBorder(screen)
}

func (table *Table) RenderBorder(screen tcell.Screen) {
	for col := table.State.Coordinates.X1; col <= table.State.Coordinates.X2; col++ {
		screen.SetContent(col, table.State.Coordinates.Y1, tcell.RuneHLine, nil, styles.StyleForView)
		screen.SetContent(col, table.State.Coordinates.Y2, tcell.RuneHLine, nil, styles.StyleForView)
	}

	for row := table.State.Coordinates.Y1; row <= table.State.Coordinates.Y2; row++ {
		screen.SetContent(table.State.Coordinates.X1, row, tcell.RuneVLine, nil, styles.StyleForView)
		screen.SetContent(table.State.Coordinates.X2, row, tcell.RuneVLine, nil, styles.StyleForView)
	}

	screen.SetContent(table.State.Coordinates.X1, table.State.Coordinates.Y1, tcell.RuneULCorner, nil, styles.StyleForView)
	screen.SetContent(table.State.Coordinates.X2, table.State.Coordinates.Y1, tcell.RuneURCorner, nil, styles.StyleForView)
	screen.SetContent(table.State.Coordinates.X1, table.State.Coordinates.Y2, tcell.RuneLLCorner, nil, styles.StyleForView)
	screen.SetContent(table.State.Coordinates.X2, table.State.Coordinates.Y2, tcell.RuneLRCorner, nil, styles.StyleForView)
}

func (table *Table) Fill(screen tcell.Screen) {
	for _, v := range table.State.Rows {
		fileView := ui.ConvertToFileView(v.File)
		var style tcell.Style
		if v.File.IsDirectory {
			style = styles.StyleForDirectory
		} else {
			style = styles.StyleForFile
		}
		ui.DrawText(screen, v.Y, table.State.X1FC, table.State.X2FC, fileView.Name, style)
		ui.DrawText(screen, v.Y, table.State.X1SC, table.State.X2SC, fileView.Info, style)
		ui.DrawText(screen, v.Y, table.State.X1TC, table.State.X2TC, fileView.UpdateTime, style)
	}
}

func (table *Table) SelectRow(screen tcell.Screen) {
	if table.State.IsActive && table.State.SelectedRow != -1 {
		selectedRow := table.State.Rows[table.State.SelectedRow]

		for col := table.State.X1FC; col <= table.State.X2TC; col++ {
			screen.SetContent(col, selectedRow.Y, ' ', nil, styles.StyleForSelectedFile)
		}
		fileView := ui.ConvertToFileView(selectedRow.File)
		ui.DrawText(screen, selectedRow.Y, table.State.X1FC, table.State.X2FC, fileView.Name, styles.StyleForSelectedFile)
		ui.DrawText(screen, selectedRow.Y, table.State.X1SC, table.State.X2SC, fileView.Info, styles.StyleForSelectedFile)
		ui.DrawText(screen, selectedRow.Y, table.State.X1TC, table.State.X2TC, fileView.UpdateTime, styles.StyleForSelectedFile)
	}
}

func Render(screen tcell.Screen, tableLeft, tableRight *Table) {
	tableLeft.RenderTable(screen)
	tableLeft.Fill(screen)

	tableRight.RenderTable(screen)
	tableRight.Fill(screen)
}
