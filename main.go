package main

import (
	"nebula-shell/shell/dock"
	"nebula-shell/shell/shared"
	"nebula-shell/svc/hyprctl"
	"nebula-shell/svc/hypripc"
	"os"
	"sync"

	"github.com/mappu/miqt/qt6"
	"github.com/mappu/miqt/qt6/mainthread"
)

func main() {
	// layershell.UseLayerShell()
	qt6.NewQApplication(os.Args)
	defer qt6.QApplication_Exec()
	for _, screen := range qt6.QGuiApplication_Screens() {
		dock.NewDock(hyprctl.HyprMonitorName(screen.Name()).Ref())
	}

	screenLock := sync.Mutex{}
	screenQueue := []*hyprctl.HyprMonitorRef{}
	timer := qt6.NewQTimer()
	timer.SetInterval(100)
	timer.OnTimerEvent(func(super func(param1 *qt6.QTimerEvent), param1 *qt6.QTimerEvent) {
		if len(screenQueue) == 0 {
			timer.Stop()
			return
		}

		screenLock.Lock()
		defer screenLock.Unlock()

		monitor := screenQueue[0]
		screenQueue = screenQueue[1:]
		screen := monitor.ToQScreen()
		if screen == nil {
			screenQueue = append(screenQueue, monitor)
			return
		}

		mainthread.Start(func() {
			dock.NewDock(monitor)
		})

	})
	shared.Ipc().EvtMonitorAdded.Add(func(ima *hypripc.IpcMonitorAdded) bool {
		screenLock.Lock()
		defer screenLock.Unlock()
		mainthread.Start(func() { timer.Start(100 * len(screenQueue)) })
		screenQueue = append(screenQueue, &ima.HyprMonitorRef)
		return false
	})
}
