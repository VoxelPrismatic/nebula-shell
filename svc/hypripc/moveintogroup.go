package hypripc

import (
	"nebula-shell/svc/hyprctl"
)

type IpcMoveIntoGroup struct {
	updates int
	hyprctl.HyprWindowRef
}

func (mon *IpcMoveIntoGroup) Update(event, value string) bool {
	mon.updates++
	switch event {
	case "moveintogroup":
		mon.Address = hyprctl.HyprWindowAddr(value)
	default:
		panic("wrong event: " + event)
	}

	return mon.updates == 1
}
