package tiles

import (
	"fmt"
	"nebula-shell/shell/corner"
	"nebula-shell/shell/qtplus"
	"nebula-shell/shell/shared"
	"nebula-shell/svc/hyprctl"
	"sync"

	"github.com/mappu/miqt/qt6"
	"github.com/mappu/miqt/qt6/mainthread"
)

type TileList struct {
	Monitor    *hyprctl.HyprMonitorRef
	Entries    []*hyprctl.HyprWindow
	Widget     *qt6.QWidget
	List       *qt6.QVBoxLayout
	Children   []*ListEntry
	ChilrenMut sync.Mutex
	Style      *qtplus.Stylesheet
}

type ListEntry struct {
	Monitor *hyprctl.HyprMonitorRef
	Parent  *qt6.QWidget
	Layout  *qt6.QVBoxLayout
	Label   *qtplus.Marquee
	Style   *qtplus.Stylesheet
	Window  *hyprctl.HyprWindow
}

var listCache = map[hyprctl.HyprMonitorName]*TileList{}

func NewListEntry(monitor *hyprctl.HyprMonitorRef, window *hyprctl.HyprWindow) *ListEntry {
	w := qt6.NewQWidget(nil)
	w.SetFixedHeight(shared.Fonts.Sans.Sizes[shared.FzText] + shared.Grid.Gap*2)
	w.SetFixedWidth(shared.Grid.InnerWidth())

	l := qt6.NewQVBoxLayout(w)
	l.SetContentsMargins(0, 0, 0, 0)

	t := qtplus.NewMarquee(nil)
	t.SetText(window.Title)
	l.AddWidget(t.QWidget)
	t.SetFixedWidth(w.Width())
	t.SetFixedHeight(w.Height())

	s := qtplus.Stylesheet{}
	s.AddTarget(w)

	s.Set("background-color", shared.Theme.Dawn.Layer.Overlay)
	s.Set("color", shared.Theme.Dawn.Text.Normal)
	s.Set("border-radius", fmt.Sprintf("%dpx", shared.Grid.Margin))

	ret := &ListEntry{
		Parent:  w,
		Label:   t,
		Layout:  l,
		Style:   &s,
		Monitor: monitor,
		Window:  window,
	}

	w.OnMousePressEvent(func(super func(event *qt6.QMouseEvent), event *qt6.QMouseEvent) {
		go hyprctl.BatchDispatch(
			[]any{"focusworkspaceoncurrentmonitor", ret.Window.Workspace.Id},
			[]any{"focuswindow", "address:" + string(ret.Window.Address)},
		)
	})

	tooltipTarget := corner.Docks[monitor.Name].Thing.ParentWidget()
	w.OnEnterEvent(func(super func(event *qt6.QEnterEvent), event *qt6.QEnterEvent) {
		qt6.QToolTip_ShowText2(event.GlobalPos(), ret.Window.Title, tooltipTarget)
		go ret.SetColors()
	})
	w.OnLeaveEvent(func(super func(event *qt6.QEvent), event *qt6.QEvent) {
		qt6.QToolTip_HideText()
		go ret.SetColors()
	})

	go ret.SetColors()
	return ret
}

func NewList(monitor *hyprctl.HyprMonitorRef) *TileList {
	if ret, ok := listCache[monitor.Name]; ok {
		return ret
	}

	w := qt6.NewQWidget(nil)
	list := &TileList{
		Monitor: monitor,
		Widget:  w,
		List:    qt6.NewQVBoxLayout(w),
		Style:   &qtplus.Stylesheet{},
	}

	list.List.SetContentsMargins(0, 0, 0, 0)
	w.SetFixedWidth(shared.Grid.InnerWidth())
	list.Style.AddTarget(w)

	listCache[monitor.Name] = list

	return list
}

func (t *TileList) SetEntries(entries []*hyprctl.HyprWindow) {
	t.ChilrenMut.Lock()
	defer t.ChilrenMut.Unlock()
	t.Entries = entries

	for i, e := range entries {
		if i >= len(t.Children) {
			mainthread.Start(func() {
				t.ChilrenMut.Lock()
				defer t.ChilrenMut.Unlock()
				c := NewListEntry(t.Monitor, e)
				t.Children = append(t.Children, c)
				t.List.AddWidget(c.Parent)
			})
		} else {
			c := t.Children[i]
			mainthread.Start(func() {
				c.Label.SetText(e.Title)
				c.Window = e
				c.Parent.SetVisible(true)
				go c.SetColors()
			})
		}
	}
	i := len(entries)
	if i < len(t.Children) {
		for _, e := range t.Children[i:] {
			mainthread.Start(func() {
				e.Parent.SetVisible(false)
				e.Label.SetPaused(true)
			})
		}
	}
}

func (e *ListEntry) GetState() shared.ActivityFlag {
	var ret shared.ActivityFlag
	mons, err := hyprctl.Monitors()
	if err != nil {
		return ret
	}

	mon := mons.Find(string(e.Monitor.Name))
	if mon == nil {
		return ret
	}

	ws, err := mon.ActiveWorkspace.Target()
	if err != nil {
		return ret
	}

	if ws.LastWindow == e.Window.Address {
		ret |= shared.EntryActive
		ret |= shared.EntryOpen
		return ret
	}

	cli, err := e.Window.Target()
	if err != nil {
		return ret
	}

	for _, m := range *mons {
		if cli.Workspace.Name == m.ActiveWorkspace.Name {
			return shared.EntryOpen
		}
	}

	return ret
}

func (e *ListEntry) SetColors() {
	e.GetState().SetColors(e.Label, e.Parent, e.Style)
}
