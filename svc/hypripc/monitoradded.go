package hypripc

import (
	"nebula-shell/svc/hyprctl"
	"strconv"
)

type IpcMonitorAdded struct {
	updates int
	hyprctl.HyprMonitorRef
	Description string
}

func (mon *IpcMonitorAdded) Update(event, value string) bool {
	mon.updates++
	switch event {
	case "monitoradded":
		break
	case "monitoraddedv2":
		parts := mustSplitN(value, 3)
		id, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		mon.Id = id
		mon.Name = parts[1]
		mon.Description = parts[2]
	default:
		panic("wrong event: " + event)
	}

	return mon.updates == 2
}
