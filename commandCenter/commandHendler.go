package commandCenter

import (
	filesystem "fileManager/fileSystem"
	"fileManager/models"
	"fileManager/ui"
	infotab "fileManager/ui/controlls/info_tab"
	tab "fileManager/ui/controlls/table"
	"fmt"
	"path/filepath"

	"github.com/gdamore/tcell/v2"
)

func HandlePollEvent(screen tcell.Screen, tableLeft, tableRight *tab.Table, infTab *infotab.InfoTab, storageLeftTable, storageRightTable []models.File) {
	for {
		ev := screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			// Выход по нажатию Escape
			if ev.Key() == tcell.KeyEscape {
				return
			}
			if ev.Key() == tcell.KeyUp {
				var table *tab.Table
				if tableLeft.State.IsActive {
					table = tableLeft
				} else if tableRight.State.IsActive {
					table = tableRight
				} else {
					continue
				}
				table.State.SelectedRow--
				selectedFile := table.State.Rows[table.State.SelectedRow]
				infTab.State.InfoCell = selectedFile.File.Path

				Render(screen, tableLeft, tableRight, infTab)
				table.SelectRow(screen)
				screen.Show()
			}
			if ev.Key() == tcell.KeyDown {
				var table *tab.Table
				if tableLeft.State.IsActive {
					table = tableLeft
				} else if tableRight.State.IsActive {
					table = tableRight
				} else {
					continue
				}
				table.State.SelectedRow++
				selectedFile := table.State.Rows[table.State.SelectedRow]
				infTab.State.InfoCell = selectedFile.File.Path

				Render(screen, tableLeft, tableRight, infTab)
				table.SelectRow(screen)
				screen.Show()
			}
			if ev.Key() == tcell.KeyLeft {
				tableRight.State.IsActive = false
				tableLeft.State.IsActive = true

				tableLeft.State.SelectedRow = 0
				infTab.State.InfoCell = tableLeft.State.CurrentTablePath

				Render(screen, tableLeft, tableRight, infTab)
				tableLeft.SelectRow(screen)
				screen.Show()
			}
			if ev.Key() == tcell.KeyRight {
				tableRight.State.IsActive = true
				tableLeft.State.IsActive = false

				tableRight.State.SelectedRow = 0
				infTab.State.InfoCell = tableLeft.State.CurrentTablePath

				Render(screen, tableLeft, tableRight, infTab)
				tableRight.SelectRow(screen)
				screen.Show()
			}
			if ev.Key() == tcell.KeyEnter {
				var table *tab.Table
				if tableLeft.State.IsActive {
					table = tableLeft
				} else if tableRight.State.IsActive {
					table = tableRight
				}
				selectedFile := table.State.Rows[table.State.SelectedRow]
				if selectedFile.File.IsDirectory {
					newFiles, _ := filesystem.GetFiles(selectedFile.File.Path)
					table.State.SetStorage(newFiles, selectedFile.File.Path)
					table.State.IsActive = true

					infTab.State.InfoCell = selectedFile.File.Path

					Render(screen, tableLeft, tableRight, infTab)
					screen.Show()
				}
			}
			if ev.Key() == tcell.KeyBackspace {
				var table *tab.Table
				if tableLeft.State.IsActive {
					table = tableLeft
				} else if tableRight.State.IsActive {
					table = tableRight
				} else {
					continue
				}
				if table.State.CurrentTablePath == rootPath {
					continue
				}
				parentDir := filepath.Dir(table.State.CurrentTablePath)
				newFiles, _ := filesystem.GetFiles(parentDir)
				table.State.SetStorage(newFiles, parentDir)
				table.State.IsActive = true

				infTab.State.InfoCell = parentDir

				Render(screen, tableLeft, tableRight, infTab)
				screen.Show()
			}
			if ev.Key() == tcell.KeyF5 {
				//copy
				var activeTable *tab.Table
				var notActiveTable *tab.Table
				if tableLeft.State.IsActive {
					activeTable = tableLeft
					notActiveTable = tableRight
				} else if tableRight.State.IsActive {
					activeTable = tableRight
					notActiveTable = tableLeft
				} else {
					continue
				}

				copiedFile := activeTable.State.Rows[activeTable.State.SelectedRow].File.Path
				folderForCopy := notActiveTable.State.CurrentTablePath

				filesystem.Copy(copiedFile, folderForCopy)

				newFiles, _ := filesystem.GetFiles(folderForCopy)
				notActiveTable.State.SetStorage(newFiles, folderForCopy)
				notActiveTable.State.IsActive = true

				infTab.State.CommandCell = fmt.Sprintf("Копирование %s в %s", copiedFile, folderForCopy)

				Render(screen, tableLeft, tableRight, infTab)
				screen.Show()
			}
			if ev.Key() == tcell.KeyF8 {
				//delete
				var activeTable *tab.Table
				if tableLeft.State.IsActive {
					activeTable = tableLeft
				} else if tableRight.State.IsActive {
					activeTable = tableRight
				} else {
					continue
				}

				delitedFile := activeTable.State.Rows[activeTable.State.SelectedRow].File

				filesystem.DeleteFile(delitedFile)

				newFiles, _ := filesystem.GetFiles(activeTable.State.CurrentTablePath)
				activeTable.State.SetStorage(newFiles, activeTable.State.CurrentTablePath)
				activeTable.State.IsActive = true

				infTab.State.CommandCell = fmt.Sprintf("Удаление %s", delitedFile.Path)

				Render(screen, tableLeft, tableRight, infTab)
				screen.Show()
			}
		case *tcell.EventResize:
			leftTableCor, rightTableCor, infTabCor := ui.GetControllsCoordinates(screen)
			tableLeft.State.SetState(leftTableCor, storageLeftTable)
			tableRight.State.SetState(rightTableCor, storageRightTable)
			infTab.State.SetState(infTabCor, "-", "-")
			Render(screen, tableLeft, tableRight, infTab)
			screen.Show()
		}
	}
}
