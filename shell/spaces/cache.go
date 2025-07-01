package spaces

import (
	"log"
	"nebula-shell/svc/hyprctl"
	"sync"
)

type wsCache map[hyprctl.HyprMonitorName]int

var (
	WsCache = wsCache{}
	wsLock  sync.Mutex
)

func (c *wsCache) Refresh() {
	wsLock.Lock()
	mons, err := hyprctl.Monitors()
	if err != nil {
		log.Fatalln(err)
	}
	for _, mon := range *mons {
		WsCache[mon.Name] = mon.ActiveWorkspace.Id
	}
	wsLock.Unlock()
}

func (c *wsCache) Restore() {
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

func (c *wsCache) Preview(t hyprctl.HyprMonitorName, ws int) {
	if !wsLock.TryLock() {
		return // do not block
	}

	batches := [][]any{}
	batches = append(batches, []any{"focusmonitor", t})
	batches = append(batches, []any{"focusworkspaceoncurrentmonitor", ws})
	_, _ = hyprctl.BatchDispatch(batches...)
	wsLock.Unlock()
}
