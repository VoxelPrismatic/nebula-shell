package hypripc

import (
	"nebula-shell/svc/hyprctl"
)

type IpcMoveOutOfGroup struct {
	updates int
	hyprctl.HyprWindowRef
}

func (mon *IpcMoveOutOfGroup) Update(event, value string) bool {
	mon.updates++
	switch event {
	case "moveoutofgroup":
		mon.Address = value
	default:
		panic("wrong event: " + event)
	}

	return mon.updates == 1
}
