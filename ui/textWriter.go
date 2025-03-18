package ui

import "github.com/gdamore/tcell/v2"

func DrawText(screen tcell.Screen, row, X1, X2 int, text string, style tcell.Style) {
	runes := []rune(text)
	for _, char := range runes {
		if X1 >= X2 {
			break
		}
		screen.SetContent(X1, row, char, nil, style)
		X1++
	}
}
