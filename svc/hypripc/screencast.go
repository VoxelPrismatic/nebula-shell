package hypripc

import (
	"strconv"
)

type CastingOwner int

const (
	CastingMonitor CastingOwner = iota
	CastingWindow
)

type IpcScreencast struct {
	updates int
	State   bool
	Owner   CastingOwner
}

func (sc *IpcScreencast) Update(event, value string) bool {
	sc.updates++
	switch event {
	case "screencast":
		parts := mustSplitN(value, 2)
		owner, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		sc.State = parts[0] == "1"
		sc.Owner = CastingOwner(owner)
	default:
		panic("wrong event: " + event)
	}

	return sc.updates == 1
}

func (sc IpcScreencast) Target() (*CastingOwner, error) {
	return &sc.Owner, nil
}
