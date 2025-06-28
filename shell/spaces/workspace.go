package spaces

import (
	"cmp"
	"fmt"
	"log"
	"nebula-shell/shell/qtplus"
	"nebula-shell/shell/shared"
	"nebula-shell/svc/hyprctl"
	"slices"
	"sync"

	"github.com/mappu/miqt/qt6"
	"github.com/mappu/miqt/qt6/mainthread"
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

type WsCache map[hyprctl.HyprMonitorName]int

var (
	wsCache = WsCache{}
	wsLock  sync.Mutex
)

func (c *WsCache) Refresh() {
	wsLock.Lock()
	mons, err := hyprctl.Monitors()
	if err != nil {
		log.Fatalln(err)
	}
	for _, mon := range *mons {
		wsCache[mon.Name] = mon.ActiveWorkspace.Id
	}
	wsLock.Unlock()
}

func (c *WsCache) Restore() {
	wsLock.Lock()

	ws, err := hyprctl.ActiveWorkspace()
	if err != nil {
		panic(err)
	}

	curMon := ws.Monitor

	batches := [][]any{}
	for mon, id := range *c {
		batches = append(batches, []any{"focusmonitor", mon})
		batches = append(batches, []any{"focusworkspaceoncurrentmonitor", id})
	}
	batches = append(batches, []any{"focusmonitor", curMon})
	_, _ = hyprctl.BatchDispatch(batches...)
	wsLock.Unlock()
}

func (c *WsCache) Preview(t hyprctl.HyprMonitorName, ws int) {
	if !wsLock.TryLock() {
		return // do not block
	}

	batches := [][]any{}
	batches = append(batches, []any{"focusmonitor", t})
	batches = append(batches, []any{"focusworkspaceoncurrentmonitor", ws})
	_, _ = hyprctl.BatchDispatch(batches...)
	wsLock.Unlock()
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
			wsCache[mon.Name] = id
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
	w.OnEnterEvent(func(super func(event *qt6.QEnterEvent), event *qt6.QEnterEvent) {
		if workspace.Target.Id != -1 {
			go wsCache.Preview(mon.Name, workspace.Target.Id)
		}
		workspace.SetColors()
	})
	w.OnLeaveEvent(func(super func(event *qt6.QEvent), event *qt6.QEvent) {
		workspace.SetColors()
	})

	workspace.SetColors()

	return workspace
}

func (e *WorkspaceEntry) SetTarget(target hyprctl.HyprWorkspaceRef) {
	e.Target = target
}

func (e *WorkspaceEntry) GetState() ActivityFlag {
	var ret ActivityFlag
	mons, err := hyprctl.Monitors()
	if err != nil {
		return ret
	}

	t, err := e.Target.Target()
	if err == nil && t.Windows == 0 {
		ret |= EntryEmpty
	}

	mon := mons.Find(string(e.Monitor.Name))
	if mon != nil && mon.ActiveWorkspace.Id == e.Target.Id {
		ret |= EntryActive
	}

	for _, mon := range *mons {
		if mon.ActiveWorkspace.Id == e.Target.Id {
			ret |= EntryOpen
			break
		}
	}

	return ret
}

type ActivityFlag int

const (
	EntryActive ActivityFlag = 1 << iota
	EntryEmpty
	EntryOpen
)

func (a ActivityFlag) All(flags ...ActivityFlag) bool {
	for _, flag := range flags {
		if a&flag == 0 {
			return false
		}
	}
	return true
}

func (a ActivityFlag) Any(flags ...ActivityFlag) bool {
	for _, flag := range flags {
		if a&flag != 0 {
			return true
		}
	}
	return false
}

func (e *WorkspaceEntry) SetColors() {
	flags := e.GetState()

	mainthread.Start(func() {
		f := e.Text.Font()
		f.SetBold(flags.Any(EntryActive, EntryEmpty))
		e.Text.SetFont(f)
		hovered := e.Widget.UnderMouse()

		switch flags {
		case EntryActive | EntryEmpty, EntryActive | EntryEmpty | EntryOpen:
			e.Stylesheet.Set("color", shared.Theme.Moon.Paint.Foam)
			if hovered {
				e.Stylesheet.Set("background-color", shared.Theme.Moon.Layer.Base)
			} else {
				e.Stylesheet.Set("background-color", shared.Theme.Dawn.Text.Normal)
			}
		case EntryActive, EntryActive | EntryOpen:
			e.Stylesheet.Set("color", shared.Theme.Dawn.Layer.Base)
			if hovered {
				e.Stylesheet.Set("background-color", shared.Theme.Moon.Layer.Base)
			} else {
				e.Stylesheet.Set("background-color", shared.Theme.Dawn.Text.Normal)
			}
		case EntryEmpty | EntryOpen:
			e.Stylesheet.Set("color", shared.Theme.Dawn.Layer.Base)
			if hovered {
				e.Stylesheet.Set("background-color", shared.Theme.Dawn.Paint.Love)
			} else {
				e.Stylesheet.Set("background-color", shared.Theme.Dawn.Paint.Rose)
			}
		case EntryEmpty:
			e.Stylesheet.Set("color", shared.Theme.Dawn.Paint.Love)
			if hovered {
				e.Stylesheet.Set("background-color", shared.Theme.Dawn.Hl.Med)
			} else {
				e.Stylesheet.Set("background-color", shared.Theme.Dawn.Layer.Overlay)
			}
		case EntryOpen:
			e.Stylesheet.Set("color", shared.Theme.Dawn.Text.Normal)
			if hovered {
				e.Stylesheet.Set("background-color", shared.Theme.Dawn.Hl.Med)
			} else {
				e.Stylesheet.Set("background-color", shared.Theme.Dawn.Hl.High)
			}
		default:
			e.Stylesheet.Set("color", shared.Theme.Dawn.Text.Normal)
			if hovered {
				e.Stylesheet.Set("background-color", shared.Theme.Dawn.Hl.Med)
			} else {
				e.Stylesheet.Set("background-color", shared.Theme.Dawn.Layer.Overlay)
			}
		}
	})
}
