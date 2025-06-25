package hypripc

import (
	"nebula-shell/svc/hyprctl"
	"strconv"
)

type IpcActiveSpecial struct {
	updates int
	hyprctl.HyprMonitorRef
	Workspace *hyprctl.HyprWorkspaceRef
}

func (mon *IpcActiveSpecial) Update(event, value string) bool {
	mon.updates++
	switch event {
	case "activespecial":
		break
	case "activespecialv2":
		parts := mustSplitN(value, 3)
		id, err := strconv.Atoi(parts[0])
		if parts[0] != "" && err != nil {
			panic(err)
		}
		if parts[0] == "" {
			mon.Workspace = nil
		} else {
			mon.Workspace = &hyprctl.HyprWorkspaceRef{Id: id, Name: parts[1]}
		}
		mon.Name = parts[2]
	default:
		panic("wrong event: " + event)
	}

	return mon.updates == 2
}
