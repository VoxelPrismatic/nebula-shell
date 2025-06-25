package hypripc

import (
	"nebula-shell/svc/hyprctl"
)

type IpcCloseWindow struct {
	updates int
	hyprctl.HyprWindowRef
}

func (win *IpcCloseWindow) Update(event, value string) bool {
	win.updates++
	switch event {
	case "closewindow":
		win.Address = value
	default:
		panic("wrong event: " + event)
	}

	return win.updates == 1
}
