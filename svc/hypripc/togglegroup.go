package hypripc

import (
	"nebula-shell/svc/hyprctl"
	"strings"
)

type IpcToggleGroup struct {
	updates  int
	Included bool
	Windows  []hyprctl.HyprWindowRef
}

func (mon *IpcToggleGroup) Update(event, value string) bool {
	mon.updates++
	switch event {
	case "togglegroup":
		parts := strings.Split(value, ",")
		mon.Windows = make([]hyprctl.HyprWindowRef, len(parts)-1)
		mon.Included = parts[0] == "1"
		for idx, addr := range parts[1:] {
			mon.Windows[idx] = hyprctl.HyprWindowRef{Address: addr}
		}
	default:
		panic("wrong event: " + event)
	}

	return mon.updates == 1
}

// Returns only the first window in the group
func (mon IpcToggleGroup) Target() (*hyprctl.HyprWindow, error) {
	return mon.Windows[0].Target()
}
