package hypripc

import (
	"fmt"
	"nebula-shell/svc/hyprctl"
)

type IpcActiveLayout struct {
	updates  int
	Keyboard string
	Layout   string
}

func (lay *IpcActiveLayout) Update(event, value string) bool {
	lay.updates++
	switch event {
	case "activelayout":
		parts := mustSplitN(value, 2)
		lay.Keyboard = parts[0]
		lay.Layout = parts[1]
	default:
		panic("wrong event: " + event)
	}

	return lay.updates == 1
}

func (lay IpcActiveLayout) Target() (*hyprctl.HyprDevKeyboard, error) {
	devices, err := hyprctl.Devices()
	if err != nil {
		return nil, err
	}
	for _, kb := range devices.Keyboards {
		if kb.Name == lay.Keyboard && kb.Layout == lay.Layout {
			return &kb, nil
		}
	}
	return nil, fmt.Errorf("keyboard not found")
}
