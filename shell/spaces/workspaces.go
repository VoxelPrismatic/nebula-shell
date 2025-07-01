package spaces

import (
	"log"
	"nebula-shell/shell/shared"
	"nebula-shell/svc/hyprctl"
	"nebula-shell/svc/hypripc"
	"sync"

	"github.com/mappu/miqt/qt6"
	"github.com/mappu/miqt/qt6/mainthread"
)

type WorkspaceGrid struct {
	LockRefresh *sync.Mutex
	LockEntry   *sync.Mutex
	Monitor     *hyprctl.HyprMonitorRef
	Entries     []*WorkspaceEntry
	Widget      *qt6.QWidget
	Grid        *qt6.QGridLayout
	Plus        *WorkspaceEntry
}

var gridCache map[hyprctl.HyprMonitorName]*WorkspaceGrid
var curWinAddr string

var _ = shared.Ipc().EvtMonitorRemoved.Add(func(imr *hypripc.IpcMonitorRemoved) bool {
	delete(gridCache, imr.Name)
	return false
})

func NewGrid(monitor *hyprctl.HyprMonitorRef) *WorkspaceGrid {
	go WsCache.Refresh()
	if gridCache == nil {
		gridCache = map[hyprctl.HyprMonitorName]*WorkspaceGrid{}
		bindRefresh()
	}
	if ret, ok := gridCache[monitor.Name]; ok {
		return ret
	}
	w := qt6.NewQWidget2()
	w.SetFixedWidth(shared.Grid.InnerWidth())
	w.SetContentsMargins(0, 0, 0, 0)
	g := &WorkspaceGrid{
		Monitor:     monitor,
		Widget:      w,
		Grid:        qt6.NewQGridLayout(w),
		LockRefresh: &sync.Mutex{},
		LockEntry:   &sync.Mutex{},
	}
	g.Grid.SetContentsMargins(0, 0, 0, 0)
	g.Grid.SetSpacing(shared.Grid.Gap)
	mainthread.Start(func() {
		g.Plus = NewEntry(g.Monitor, -1, hyprctl.HyprWorkspaceRef{Id: -1}, g)
	})

	w.OnLeaveEvent(func(super func(event *qt6.QEvent), event *qt6.QEvent) {
		if w.UnderMouse() {
			return
		}
		go WsCache.Restore()
	})
	w.OnEnterEvent(func(super func(event *qt6.QEnterEvent), event *qt6.QEnterEvent) {
		go func() {
			go WsCache.Refresh()
			win, err := hyprctl.ActiveWindow()
			if err != nil {
				log.Fatalln(err)
			}
			curWinAddr = string(win.Address)
		}()
	})

	g.Refresh(nil)

	gridCache[monitor.Name] = g
	return g
}

func propRefresh() {
	wss, err := hyprctl.Workspaces()
	if err != nil {
		return
	}
	for _, g := range gridCache {
		go g.Refresh(wss)
	}
}

func bindRefresh() {
	bounceTimer := qt6.NewQTimer()
	bounceTimer.SetSingleShot(true)
	bounceTimer.OnTimeout(propRefresh)

	bounce := func() {
		go mainthread.Start(func() { bounceTimer.Start(10) })
	}

	shared.Ipc().EvtDestroyWorkspace.Add(func(idw *hypripc.IpcDestroyWorkspace) bool {
		bounce()
		return false
	})

	shared.Ipc().EvtDestroyWorkspace.Add(func(idw *hypripc.IpcDestroyWorkspace) bool {
		bounce()
		return false
	})

	shared.Ipc().EvtWorkspace.Add(func(iw *hypripc.IpcWorkspace) bool {
		bounce()
		return false
	})
}

func (g *WorkspaceGrid) Refresh(wss *[]hyprctl.HyprWorkspace) {
	if !g.LockRefresh.TryLock() {
		return // Refresh is already in progress
	}
	defer g.LockRefresh.Unlock()

	if wss == nil {
		var err error
		wss, err = hyprctl.Workspaces()
		if err != nil {
			panic(err)
		}
		if wss == nil {
			log.Fatalf("nil workspaces")
		}
	}

	empty := false
	for i, ws := range *wss {
		var e *WorkspaceEntry
		if i >= len(g.Entries) {
			mainthread.Start(func() {
				e = NewEntry(g.Monitor, i, ws.HyprWorkspaceRef, g)
				g.LockEntry.Lock()
				g.Entries = append(g.Entries, e)
				g.LockEntry.Unlock()
				g.Grid.AddWidget2(e.Widget, (i/3)*3, i%3)
			})
		} else {
			e = g.Entries[i]
			e.SetTarget(ws.HyprWorkspaceRef)
			mainthread.Start(func() { e.Widget.SetVisible(true) })
			go e.SetColors()
		}
		empty = empty || ws.Windows == 0 && ws.Monitor == g.Monitor.Name
	}

	i := len(*wss)
	if i < len(g.Entries) {
		g.LockEntry.Lock()
		for _, e := range g.Entries[i:] {
			mainthread.Start(func() { e.Widget.SetVisible(false) })
		}
		g.LockEntry.Unlock()
	}
	mainthread.Start(func() {
		g.Grid.AddWidget2(g.Plus.Widget, (i/3)*3, i%3)
		g.Plus.Widget.SetVisible(!empty)
	})
}
