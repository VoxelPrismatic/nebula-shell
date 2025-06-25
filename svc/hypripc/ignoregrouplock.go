package hypripc

type IpcIgnoreGroupLock struct {
	updates int
	State   bool
}

func (fs *IpcIgnoreGroupLock) Update(event, value string) bool {
	fs.updates++
	switch event {
	case "ignoregrouplock":
		fs.State = value == "1"
	default:
		panic("wrong event: " + event)
	}

	return fs.updates == 1
}

func (fs IpcIgnoreGroupLock) Target() (*bool, error) {
	return &fs.State, nil
}
