package hypripc

import (
	"nebula-shell/svc/hyprctl"
)

type IpcFloatChanged struct {
	updates int
	hyprctl.HyprWindowRef
	Floating bool
}

func (win *IpcFloatChanged) Update(event, value string) bool {
	win.updates++
	switch event {
	case "changefloatingmode":
		parts := mustSplitN(value, 2)
		win.Address = parts[0]
		win.Floating = parts[1] == "1"
	default:
		panic("wrong event: " + event)
	}

	return win.updates == 1
}
