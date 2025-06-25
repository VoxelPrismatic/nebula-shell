package hypripc

import (
	"nebula-shell/svc/hyprctl"
)

type IpcOpenWindow struct {
	updates int
	hyprctl.HyprWindowRef
	Class     string
	Workspace hyprctl.HyprWorkspaceRef
}

func (win *IpcOpenWindow) Update(event, value string) bool {
	win.updates++
	switch event {
	case "openwindow":
		parts := mustSplitN(value, 4)
		win.Address = parts[0]
		win.Workspace.Name = parts[1]
		win.Class = parts[2]
		win.Title = parts[3]
	default:
		panic("wrong event: " + event)
	}

	return win.updates == 1
}
