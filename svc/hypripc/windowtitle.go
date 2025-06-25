package hypripc

import (
	"nebula-shell/svc/hyprctl"
)

type IpcWindowTitle struct {
	updates int
	hyprctl.HyprWindowRef
}

func (win *IpcWindowTitle) Update(event, value string) bool {
	win.updates++
	switch event {
	case "windowtitle":
		break
	case "windowtitlev2":
		parts := mustSplitN(value, 2)
		win.Address = parts[0]
		win.Title = parts[1]
	default:
		panic("wrong event: " + event)
	}

	return win.updates == 2
}
