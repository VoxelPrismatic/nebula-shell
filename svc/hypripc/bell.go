package hypripc

import (
	"nebula-shell/svc/hyprctl"
)

type IpcBell struct {
	updates int
	hyprctl.HyprWindowRef
}

func (win *IpcBell) Update(event, value string) bool {
	win.updates++
	switch event {
	case "bell":
		win.Address = hyprctl.HyprWindowAddr(value)
	default:
		panic("wrong event: " + event)
	}

	return win.updates == 1
}
