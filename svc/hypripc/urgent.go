package hypripc

import (
	"nebula-shell/svc/hyprctl"
)

type IpcUrgent struct {
	updates int
	hyprctl.HyprWindowRef
}

func (win *IpcUrgent) Update(event, value string) bool {
	win.updates++
	switch event {
	case "changefloatingmode":
		win.Address = value
	default:
		panic("wrong event: " + event)
	}

	return win.updates == 1
}
