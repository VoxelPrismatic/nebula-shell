package tiles

import (
	"cmp"
	"log"
	"nebula-shell/shell/shared"
	"nebula-shell/svc/desktop"
	"nebula-shell/svc/hyprctl"
	"nebula-shell/svc/hypripc"
	"slices"
	"sync"

	"github.com/mappu/miqt/qt6"
	"github.com/mappu/miqt/qt6/mainthread"
)

type TileGrid struct {
	LockRefresh *sync.Mutex
	LockEntry   *sync.Mutex
	Monitor     *hyprctl.HyprMonitorRef
	Entries     []*TileEntry
	Widget      *qt6.QWidget
	Grid        *qt6.QGridLayout
	Plus        *TileEntry
}

var gridCache map[hyprctl.HyprMonitorName]*TileGrid

var _ = shared.Ipc().EvtMonitorRemoved.Add(func(imr *hypripc.IpcMonitorRemoved) bool {
	delete(gridCache, imr.Name)
	delete(listCache, imr.Name)
	return false
})

func NewGrid(monitor *hyprctl.HyprMonitorRef) *TileGrid {
	if gridCache == nil {
		gridCache = map[hyprctl.HyprMonitorName]*TileGrid{}
		bindRefresh()
	}
	if ret, ok := gridCache[monitor.Name]; ok {
		return ret
	}
	w := qt6.NewQWidget2()
	w.SetFixedWidth(shared.Grid.InnerWidth())
	g := &TileGrid{
		Monitor:     monitor,
		Widget:      w,
		Grid:        qt6.NewQGridLayout(w),
		LockRefresh: &sync.Mutex{},
		LockEntry:   &sync.Mutex{},
	}
	g.Grid.SetContentsMargins(0, 0, 0, 0)
	g.Grid.SetSpacing(shared.Grid.Gap)
	// mainthread.Start(func() {
	// g.Plus = NewEntry(g.Monitor, -1, hyprctl.HyprWorkspaceRef{Id: -1}, g)
	// })

	g.Refresh(nil)

	gridCache[monitor.Name] = g
	return g
}

type propData map[*desktop.DesktopFilePlus][]*hyprctl.HyprWindow

func propRefreshData() *propData {
	wins, err := hyprctl.Clients()
	if err != nil {
		return nil
	}
	prop := propData{}
	exist := map[string]*desktop.DesktopFilePlus{}
	for _, win := range *wins {
		df := win.DesktopFile()
		if e, ok := exist[df.FilePath_]; ok {
			df = e
		} else {
			exist[df.FilePath_] = df
		}

		prop[df] = append(prop[df], &win)
	}

	return &prop
}
func propRefresh() {
	prop := propRefreshData()
	for _, g := range gridCache {
		go g.Refresh(prop)
	}
}

func bindRefresh() {
	bounceTimer := qt6.NewQTimer()
	bounceTimer.SetSingleShot(true)
	bounceTimer.OnTimeout(propRefresh)

	bounce := func() {
		mainthread.Start(func() { bounceTimer.Start(50) })
	}

	shared.Ipc().EvtOpenWindow.Add(func(idw *hypripc.IpcOpenWindow) bool {
		bounce()
		return false
	})

	shared.Ipc().EvtCloseWindow.Add(func(idw *hypripc.IpcCloseWindow) bool {
		bounce()
		return false
	})

	shared.Ipc().EvtActiveWindow.Add(func(iw *hypripc.IpcActiveWindow) bool {
		bounce()
		return false
	})
}

func (g *TileGrid) Refresh(prop *propData) {
	if !g.LockRefresh.TryLock() {
		return // Refresh is already in progress
	}
	defer g.LockRefresh.Unlock()

	if prop == nil {
		prop = propRefreshData()
		if prop == nil {
			log.Fatalf("nil clients")
		}
	}

	apps := make([]*desktop.DesktopFilePlus, len(*prop))
	i := 0
	for key := range *prop {
		apps[i] = key
		i++
	}
	slices.SortFunc(apps, func(a, b *desktop.DesktopFilePlus) int {
		return cmp.Compare(a.Name, b.Name)
	})

	for i, key := range apps {
		arr := (*prop)[key]
		slices.SortFunc(arr, func(a, b *hyprctl.HyprWindow) int {
			return cmp.Compare(a.Pid, b.Pid)
		})

		if i >= len(g.Entries) {
			mainthread.Start(func() {
				e := NewEntry(g.Monitor, key, arr, g)
				g.LockEntry.Lock()
				g.Entries = append(g.Entries, e)
				g.LockEntry.Unlock()
				i := g.Grid.Count()
				g.Grid.AddWidget2(e.Widget, (i/3)*3, i%3)
			})
		} else {
			e := g.Entries[i]
			e.SetTarget(key, arr)
			mainthread.Start(func() { e.Widget.SetVisible(true) })
			go e.SetColors()
		}
	}

	i = len(apps)
	if i < len(g.Entries) {
		g.LockEntry.Lock()
		for _, e := range g.Entries[i:] {
			mainthread.Start(func() { e.Widget.SetVisible(false) })
		}
		g.LockEntry.Unlock()
	}
	// mainthread.Start(func() {
	// g.Grid.AddWidget2(g.Plus.Widget, (i/3)*3, i%3)
	// }
}
