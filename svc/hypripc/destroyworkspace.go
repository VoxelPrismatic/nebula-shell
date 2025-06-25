package hypripc

import (
	"nebula-shell/svc/hyprctl"
	"strconv"
)

type IpcDestroyWorkspace struct {
	updates int
	hyprctl.HyprWorkspaceRef
}

func (ws *IpcDestroyWorkspace) Update(event, value string) bool {
	ws.updates++
	switch event {
	case "destroyworkspace":
		break
	case "destroyworkspacev2":
		parts := mustSplitN(value, 2)
		id, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		ws.Id = id
		ws.Name = parts[1]
	default:
		panic("wrong event: " + event)
	}

	return ws.updates == 2
}
