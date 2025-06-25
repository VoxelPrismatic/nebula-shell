package hypripc

import (
	"nebula-shell/svc/hyprctl"
)

type IpcActiveWindow struct {
	updates int
	Class   string
	hyprctl.HyprWindowRef
}

func (win *IpcActiveWindow) Update(event, value string) bool {
	win.updates++
	switch event {
	case "activewindow":
		parts := mustSplitN(value, 2)
		win.Class = parts[0]
		win.Title = parts[1]
	case "activewindowv2":
		win.Address = value
	default:
		panic("wrong event: " + event)
	}

	return win.updates == 2
}
