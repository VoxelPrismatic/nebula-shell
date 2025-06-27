package hypripc

import (
	"nebula-shell/svc/hyprctl"
)

type IpcCloseLayer struct {
	updates int
	hyprctl.HyprLayerRef
}

func (lay *IpcCloseLayer) Update(event, value string) bool {
	lay.updates++
	switch event {
	case "closelayer":
		lay.Namespace = value
	default:
		panic("wrong event: " + event)
	}

	return lay.updates == 1
}
