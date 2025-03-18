package ui

import "github.com/gdamore/tcell/v2"

type CoordinatesControl struct {
	X1, X2, Y1, Y2 int
}

func GetControllsCoordinates(screen tcell.Screen) (leftTableCor, rightTableCor, infoTabCor CoordinatesControl) {
	width, heigh := screen.Size()

	infoTabCor = CoordinatesControl{}
	leftTableCor = CoordinatesControl{}
	rightTableCor = CoordinatesControl{}

	infoTabCor.X1 = 0
	infoTabCor.X2 = width - 1
	infoTabCor.Y1 = heigh - 3
	infoTabCor.Y2 = heigh - 1

	leftTableCor.X1 = 0
	leftTableCor.X2 = width / 2
	leftTableCor.Y1 = 0
	leftTableCor.Y2 = heigh - 5

	rightTableCor.X1 = leftTableCor.X2 + 2
	rightTableCor.X2 = width - 1
	rightTableCor.Y1 = 0
	rightTableCor.Y2 = heigh - 5

	return leftTableCor, rightTableCor, infoTabCor
}
