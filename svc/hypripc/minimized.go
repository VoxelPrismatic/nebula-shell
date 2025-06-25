package hypripc

import (
	"nebula-shell/svc/hyprctl"
)

type IpcMinimized struct {
	updates int
	hyprctl.HyprWindowRef
	State bool
}

func (win *IpcMinimized) Update(event, value string) bool {
	win.updates++
	switch event {
	case "minimized":
		parts := mustSplitN(value, 2)
		win.Address = parts[0]
		win.State = parts[1] == "1"
	default:
		panic("wrong event: " + event)
	}

	return win.updates == 1
}
