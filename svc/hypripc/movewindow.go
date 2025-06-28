package hypripc

import (
	"nebula-shell/svc/hyprctl"
	"strconv"
)

type IpcMoveWindow struct {
	updates int
	hyprctl.HyprWindowRef
	Workspace hyprctl.HyprWorkspaceRef
}

func (win *IpcMoveWindow) Update(event, value string) bool {
	win.updates++
	switch event {
	case "movewindow":
		break
	case "movewindowv2":
		parts := mustSplitN(value, 3)
		id, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		win.Address = hyprctl.HyprWindowAddr(parts[0])
		win.Workspace.Id = id
		win.Workspace.Name = parts[2]
	default:
		panic("wrong event: " + event)
	}

	return win.updates == 2
}
