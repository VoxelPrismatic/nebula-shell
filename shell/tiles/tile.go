package tiles

import (
	"nebula-shell/shell/corner"
	"nebula-shell/shell/qtplus"
	"nebula-shell/shell/shared"
	"nebula-shell/shell/spaces"
	"nebula-shell/svc/desktop"
	"nebula-shell/svc/hyprctl"

	"github.com/mappu/miqt/qt6"
)

type TileEntry struct {
	Target     []*hyprctl.HyprWindow
	ptr        int
	Widget     *qt6.QWidget
	Monitor    *hyprctl.HyprMonitorRef
	Stylesheet *qtplus.Stylesheet
	Text       *qt6.QLabel
	Desktop    *desktop.DesktopFilePlus
	Parent     *TileGrid
}

func NewEntry(mon *hyprctl.HyprMonitorRef, key *desktop.DesktopFilePlus, ref []*hyprctl.HyprWindow, parent *TileGrid) *TileEntry {
	workspace := &TileEntry{
		Target:     ref,
		Widget:     qt6.NewQWidget2(),
		Monitor:    mon,
		Stylesheet: &qtplus.Stylesheet{},
		Text:       qt6.NewQLabel2(),
		Desktop:    key,
		Parent:     parent,
	}

	w := workspace.Widget
	w.SetContentsMargins(0, 0, 0, 0)
	s := workspace.Stylesheet
	t := workspace.Text
	s.AddTarget(w)
	s.Set("border-radius", "4px")
	s.Set("background-color", shared.Theme.Dawn.Layer.Overlay)
	sz := shared.Grid.CellWidth()
	w.SetFixedSize2(sz, sz)

	t.SetFont(shared.Fonts.Sans.ToQFont(6))
	t.SetWordWrap(true)
	if workspace.Target == nil {
		t.SetText("+")
	} else {
		t.SetPixmap(qt6.QIcon_FromTheme(key.Icon).Pixmap2(24, 24))
	}

	layout := qt6.NewQVBoxLayout(w)
	layout.AddWidget(workspace.Text.QWidget)
	layout.SetContentsMargins(0, 0, 0, 0)
	workspace.Text.SetAlignment(qt6.AlignCenter)

	w.OnMousePressEvent(func(super func(event *qt6.QMouseEvent), event *qt6.QMouseEvent) {
		workspace.ptr = (workspace.ptr + 1) % len(workspace.Target)
		target := workspace.Target[workspace.ptr]
		go func() {
			_, _ = hyprctl.BatchDispatch(
				[]any{"focusworkspaceoncurrentmonitor", target.Workspace.Id},
				[]any{"focuswindow", "address:" + string(target.Address)},
			)
			spaces.WsCache.Refresh()
		}()
	})

	tooltipTarget := corner.Docks[mon.Name].Thing.ParentWidget()
	w.OnEnterEvent(func(super func(event *qt6.QEnterEvent), event *qt6.QEnterEvent) {
		qt6.QToolTip_ShowText2(event.GlobalPos(), workspace.Desktop.Name, tooltipTarget)
		go workspace.SetColors()
	})
	w.OnLeaveEvent(func(super func(event *qt6.QEvent), event *qt6.QEvent) {
		qt6.QToolTip_HideText()
		go workspace.SetColors()
	})

	workspace.SetColors()

	return workspace
}

func (e *TileEntry) SetTarget(df *desktop.DesktopFilePlus, target []*hyprctl.HyprWindow) {
	e.Target = target
	e.Desktop = df
	if df != nil {
		pix := shared.IconCache.Get(df.Icon, 24)
		e.Text.SetPixmap(pix)
	} else {
		pix := shared.IconCache.Get("folder-important", 24)
		e.Text.SetPixmap(pix)
	}
	go e.SetColors()
}

func (e *TileEntry) GetState() shared.ActivityFlag {
	var ret shared.ActivityFlag
	mons, err := hyprctl.Monitors()
	if err != nil {
		return ret
	}

	mon := mons.Find(string(e.Monitor.Name))
	if mon != nil {
		ws, err := mon.ActiveWorkspace.Target()
		if err != nil {
			return ret
		}
		for _, win := range e.Target {
			if win.Address == ws.LastWindow {
				ret |= shared.EntryActive
				break
			}
		}
	}

	return ret
}

func (e *TileEntry) SetColors() {
	flags := e.GetState()
	flags.SetColors(e.Text, e.Widget, e.Stylesheet)

	if flags.Any(shared.EntryActive) {
		if list, ok := listCache[e.Monitor.Name]; ok {
			list.SetEntries(e.Target)
		}
	}
}
