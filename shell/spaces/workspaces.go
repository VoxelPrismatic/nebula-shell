package spaces

import (
	"nebula-shell/svc/hyprctl"

	"github.com/mappu/miqt/qt6"
)

type WorkspaceGrid struct {
	Monitor *qt6.QScreen
	Entries []WorkspaceEntry
	Widget  *qt6.QWidget
	Grid    *qt6.QGridLayout
}

var gridCache = map[string]*WorkspaceGrid{}

func NewGrid(screen *qt6.QScreen) *WorkspaceGrid {
	if ret, ok := gridCache[screen.Name()]; ok {
		return ret
	}
	w := qt6.NewQWidget2()
	g := &WorkspaceGrid{
		Monitor: screen,
		Widget:  w,
		Grid:    qt6.NewQGridLayout(w),
	}

	wss, err := hyprctl.Workspaces()
	if err != nil {
		panic(err)
	}

	for _, ws := range *wss {
		g.Grid.AddWidget(NewEntry(screen, ws.HyprWorkspaceRef).Widget)
	}

	gridCache[screen.Name()] = g
	return g
}
