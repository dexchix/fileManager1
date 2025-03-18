package info_tab

import "fileManager/ui"

type InfoTabState struct {
	Coordinates ui.CoordinatesControl

	X1da, X2da, Y1da, Y2da int

	X1FC, X2FC int

	X1SC, X2SC int

	InfoCell, CommandCell string
}

func (infoTabState *InfoTabState) SetState(Coordinates ui.CoordinatesControl, infoCell, commandCell string) {
	infoTabState.Coordinates = Coordinates

	infoTabState.X1da = infoTabState.Coordinates.X1 + 1
	infoTabState.Y1da = infoTabState.Coordinates.Y1 + 1
	infoTabState.X2da = infoTabState.Coordinates.X2 - 1
	infoTabState.Y2da = infoTabState.Coordinates.Y2 - 1

	www := (Coordinates.X2-Coordinates.X1)/2 - 2
	infoTabState.X1FC = infoTabState.X1da
	infoTabState.X2FC = infoTabState.X1FC + www

	infoTabState.X1SC = infoTabState.X2FC + 2
	infoTabState.X2SC = infoTabState.X2da

	infoTabState.InfoCell = infoCell
	infoTabState.CommandCell = commandCell
}
