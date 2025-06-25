package hypripc

import (
	"nebula-shell/svc/hyprctl"
	"strconv"
)

type IpcRenameWorkspace struct {
	updates int
	hyprctl.HyprWorkspaceRef
}

func (mon *IpcRenameWorkspace) Update(event, value string) bool {
	mon.updates++
	switch event {
	case "moveworkspace":
		break
	case "moveworkspacev2":
		parts := mustSplitN(value, 2)
		id, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		mon.Id = id
		mon.Name = parts[1]
	default:
		panic("wrong event: " + event)
	}

	return mon.updates == 2
}
