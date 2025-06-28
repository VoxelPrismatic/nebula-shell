package hypripc

import (
	"nebula-shell/svc/hyprctl"
	"strconv"
)

type IpcMoveWorkspace struct {
	updates int
	hyprctl.HyprMonitorRef
	Workspace hyprctl.HyprWorkspaceRef
}

func (mon *IpcMoveWorkspace) Update(event, value string) bool {
	mon.updates++
	switch event {
	case "moveworkspace":
		break
	case "moveworkspacev2":
		parts := mustSplitN(value, 3)
		id, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		mon.Workspace.Id = id
		mon.Workspace.Name = parts[1]
		mon.Name = hyprctl.HyprMonitorName(parts[2])
	default:
		panic("wrong event: " + event)
	}

	return mon.updates == 2
}
