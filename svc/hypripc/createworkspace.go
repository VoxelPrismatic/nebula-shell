package hypripc

import (
	"nebula-shell/svc/hyprctl"
	"strconv"
)

type IpcCreateWorkspace struct {
	updates int
	hyprctl.HyprWorkspaceRef
}

func (ws *IpcCreateWorkspace) Update(event, value string) bool {
	ws.updates++
	switch event {
	case "createworkspace":
		break
	case "createworkspacev2":
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
