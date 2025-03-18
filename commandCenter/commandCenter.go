package commandCenter

import (
	filesystem "fileManager/fileSystem"
	"fileManager/models"
	"fileManager/ui"
	infotab "fileManager/ui/controlls/info_tab"
	"fileManager/ui/controlls/table"

	"github.com/gdamore/tcell/v2"
)

var rootPath string = "C:\\"
var leftTableStorage []models.File
var rightTableStorage []models.File

var leftTable table.Table
var rightTable table.Table
var infTab infotab.InfoTab

var screen tcell.Screen

func InitializeUI() {
	Render(screen, &leftTable, &rightTable, &infTab)
	screen.Show()

	HandlePollEvent(screen, &leftTable, &rightTable, &infTab, leftTableStorage, rightTableStorage)
}

func Render(screen tcell.Screen, leftTable, rightTable *table.Table, infoTab *infotab.InfoTab) {
	screen.Clear()
	table.Render(screen, leftTable, rightTable)
	infTab.RenderInfoTab(screen)
}

func init() {
	screen, _ = tcell.NewScreen()
	screen.Init()

	leftTableStorage, _ = filesystem.GetFiles(rootPath)
	rightTableStorage, _ = filesystem.GetFiles(rootPath)

	leftTableCor, rightTableCor, infoTabCor := ui.GetControllsCoordinates(screen)

	stateForLeftTable := table.TableState{}
	stateForLeftTable.SetState(leftTableCor, leftTableStorage)
	stateForLeftTable.CurrentTablePath = rootPath
	leftTable = table.Table{State: &stateForLeftTable}

	stateForRightTable := table.TableState{}
	stateForRightTable.SetState(rightTableCor, rightTableStorage)
	stateForLeftTable.CurrentTablePath = rootPath
	rightTable = table.Table{State: &stateForRightTable}

	stateForInfoTab := infotab.InfoTabState{}
	stateForInfoTab.SetState(infoTabCor, rootPath, "-")
	infTab = infotab.InfoTab{State: &stateForInfoTab}
}
