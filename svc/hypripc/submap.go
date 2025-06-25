package hypripc

type IpcSubmap struct {
	updates int
	Name    string
	Default bool
}

func (sm *IpcSubmap) Update(event, value string) bool {
	sm.updates++
	switch event {
	case "submap":
		sm.Name = value
		sm.Default = value == ""
	default:
		panic("wrong event: " + event)
	}

	return sm.updates == 1
}

func (sm IpcSubmap) Target() (*string, error) {
	return &sm.Name, nil
}
