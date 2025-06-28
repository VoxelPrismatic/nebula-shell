package hypripc

import (
	"nebula-shell/svc/hyprctl"
)

type IpcWindowTitle struct {
	updates int
	hyprctl.HyprWindowRef
}

func (obj *IpcWindowTitle) Update(event, value string) bool {
	obj.updates++
	switch event {
	case "windowtitle":
		break
	case "windowtitlev2":
		parts := mustSplitN(value, 2)
		obj.Address = hyprctl.HyprWindowAddr(parts[0])
		obj.Title = parts[1]
	default:
		panic("wrong event: " + event)
	}

	return obj.updates == 2
}
