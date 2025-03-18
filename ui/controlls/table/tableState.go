package table

import (
	"fileManager/models"
	"fileManager/ui"
)

type TableState struct {
	Coordinates *ui.CoordinatesControl

	X1da, X2da, Y1da, Y2da int

	X1FC, X2FC int

	X1SC, X2SC int

	X1TC, X2TC int

	Rows map[int]Row

	SelectedRow int

	StorageFiles []models.File

	IsActive bool

	CurrentTablePath string
}

func (tableState *TableState) SetState(coordinates ui.CoordinatesControl, files []models.File) {
	tableState.Coordinates = &coordinates

	tableState.X1da = tableState.Coordinates.X1 + 1
	tableState.Y1da = tableState.Coordinates.Y1 + 1
	tableState.X2da = tableState.Coordinates.X2 - 1
	tableState.Y2da = tableState.Coordinates.Y2 - 1

	totalWidth := tableState.X2da - tableState.X1da
	separatorWidth := 2
	columnWidth := (totalWidth - 2*separatorWidth) / 3

	tableState.X1FC = tableState.X1da
	tableState.X2FC = tableState.X1FC + columnWidth

	tableState.X1SC = tableState.X2FC + separatorWidth
	tableState.X2SC = tableState.X1SC + columnWidth

	tableState.X1TC = tableState.X2SC + separatorWidth
	tableState.X2TC = tableState.X1TC + columnWidth

	tableState.SetStorage(files, tableState.CurrentTablePath)
}

func (tableState *TableState) SetStorage(files []models.File, path string) {
	tableState.CurrentTablePath = path
	tableState.StorageFiles = files
	tableState.Rows = make(map[int]Row)
	heighMax := tableState.Y2da - tableState.Y1da
	counterY := tableState.Y1da

	for c, v := range files {
		if counterY <= heighMax+1 {
			tableState.Rows[c] = Row{Y: counterY, File: v}
			counterY++
		}
	}

	tableState.IsActive = false
	tableState.SelectedRow = -1
}
