package hypripc

import (
	"nebula-shell/svc/hyprctl"
)

type IpcOpenLayer struct {
	updates int
	hyprctl.HyprLayerRef
}

func (lay *IpcOpenLayer) Update(event, value string) bool {
	lay.updates++
	switch event {
	case "openlayer":
		lay.Namespace = value
	default:
		panic("wrong event: " + event)
	}

	return lay.updates == 1
}
