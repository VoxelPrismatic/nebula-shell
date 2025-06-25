package hypripc

import (
	"nebula-shell/svc/hyprctl"
	"strconv"
	"strings"
)

type IpcWorkspace struct {
	updates int
	hyprctl.HyprWorkspaceRef
}

func (ws *IpcWorkspace) Update(event, value string) bool {
	ws.updates++
	switch event {
	case "workspace":
		break
	case "workspacev2":
		parts := strings.SplitN(value, ",", 2)
		if len(parts) != 2 {
			panic("ipc: spec changed")
		}
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
