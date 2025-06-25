package hypripc

import (
	"nebula-shell/svc/hyprctl"
	"strconv"
)

type IpcFocusedMonitor struct {
	updates int
	hyprctl.HyprMonitorRef
	Workspace hyprctl.HyprWorkspaceRef
}

func (mon *IpcFocusedMonitor) Update(event, value string) bool {
	mon.updates++
	switch event {
	case "focusedmon":
		parts := mustSplitN(value, 2)
		mon.Name = parts[0]
		mon.Workspace.Name = parts[1]
	case "focusedmonv2":
		parts := mustSplitN(value, 2)
		id, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		mon.Name = parts[0]
		mon.Workspace.Id = id
	default:
		panic("wrong event: " + event)
	}

	return mon.updates == 2
}
