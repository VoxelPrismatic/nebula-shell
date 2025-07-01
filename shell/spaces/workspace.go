package spaces

import (
	"cmp"
	"fmt"
	"nebula-shell/shell/qtplus"
	"nebula-shell/shell/shared"
	"nebula-shell/svc/hyprctl"
	"slices"

	"github.com/mappu/miqt/qt6"
)

var SpaceNames = []string{
	"α", "β", "δ",
	"ζ", "ξ", "ϟ",
	"λ", "π", "μ",
	"τ", "ω", "ϰ",
}

type WorkspaceEntry struct {
	Target     hyprctl.HyprWorkspaceRef
	Widget     *qt6.QWidget
	Monitor    *hyprctl.HyprMonitorRef
	Stylesheet *qtplus.Stylesheet
	Text       *qt6.QLabel
	Idx        int
	Parent     *WorkspaceGrid
}

func getMaxWsId() int {
	wss, err := hyprctl.Workspaces()
	if err != nil {
		return 100
	}

	return slices.MaxFunc(*wss, func(a, b hyprctl.HyprWorkspace) int {
		return cmp.Compare(a.Id, b.Id)
	}).Id + 1
}

func NewEntry(mon *hyprctl.HyprMonitorRef, idx int, ref hyprctl.HyprWorkspaceRef, parent *WorkspaceGrid) *WorkspaceEntry {
	workspace := &WorkspaceEntry{
		Target:     ref,
		Widget:     qt6.NewQWidget2(),
		Monitor:    mon,
		Stylesheet: &qtplus.Stylesheet{},
		Text:       qt6.NewQLabel2(),
		Idx:        idx,
		Parent:     parent,
	}

	w := workspace.Widget
	s := workspace.Stylesheet
	t := workspace.Text
	s.AddTarget(w)
	s.Set("border-radius", "4px")
	s.Set("background-color", shared.Theme.Dawn.Layer.Overlay)
	sz := shared.Grid.CellWidth()
	w.SetFixedSize2(sz, sz)

	t.SetFont(shared.Fonts.Sans.Preset(shared.FzHeader))
	if workspace.Target.Id == -1 {
		t.SetText("+")
	} else if idx >= len(SpaceNames) {
		t.SetText(fmt.Sprint(idx + 1))
	} else {
		t.SetText(SpaceNames[idx])
	}

	layout := qt6.NewQVBoxLayout(w)
	layout.AddWidget(workspace.Text.QWidget)
	workspace.Text.SetAlignment(qt6.AlignCenter)

	w.OnMousePressEvent(func(super func(event *qt6.QMouseEvent), event *qt6.QMouseEvent) {
		id := workspace.Target.Id
		if id == -1 {
			id = getMaxWsId()
		}
		batches := [][]any{}
		switch event.Button() {
		case qt6.RightButton:
			batches = append(batches, []any{"focuswindow", "address:" + curWinAddr})
			batches = append(batches, []any{"movetoworkspace", id})
			fallthrough
		case qt6.LeftButton:
			wsLock.Lock()
			for key, val := range WsCache {
				if val == id {
					WsCache[key] = WsCache[mon.Name]
					break
				}
			}
			WsCache[mon.Name] = id
			wsLock.Unlock()
			batches = append(batches, []any{"focusworkspaceoncurrentmonitor", id})
		case qt6.MiddleButton:
			batches = append(batches, []any{"focuswindow", "address:" + curWinAddr})
			batches = append(batches, []any{"movetoworkspacesilent", id})
		}
		go func() {
			_, _ = hyprctl.BatchDispatch(batches...)
		}()
		// workspace.Parent.Refresh(nil)
	})

	w.SetMouseTracking(true)
	w.OnEnterEvent(func(super func(event *qt6.QEnterEvent), event *qt6.QEnterEvent) {
		go workspace.SetColors()
		if event.Buttons()&qt6.LeftButton == 0 {
			return
		}
		if workspace.Target.Id != -1 {
			go WsCache.Preview(mon.Name, workspace.Target.Id)
		}
	})
	w.OnLeaveEvent(func(super func(event *qt6.QEvent), event *qt6.QEvent) {
		go workspace.SetColors()
	})

	workspace.SetColors()

	return workspace
}

func (e *WorkspaceEntry) SetTarget(target hyprctl.HyprWorkspaceRef) {
	e.Target = target
}

func (e *WorkspaceEntry) GetState() shared.ActivityFlag {
	var ret shared.ActivityFlag
	mons, err := hyprctl.Monitors()
	if err != nil {
		return ret
	}

	t, err := e.Target.Target()
	if err == nil && t.Windows == 0 {
		ret |= shared.EntryEmpty
	}

	mon := mons.Find(string(e.Monitor.Name))
	if mon != nil && mon.ActiveWorkspace.Id == e.Target.Id {
		ret |= shared.EntryActive
	}

	for _, mon := range *mons {
		if mon.ActiveWorkspace.Id == e.Target.Id {
			ret |= shared.EntryOpen
			break
		}
	}

	return ret
}

func (e *WorkspaceEntry) SetColors() {
	e.GetState().SetColors(e.Text, e.Widget, e.Stylesheet)
}
