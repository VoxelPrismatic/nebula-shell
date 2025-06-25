package hypripc

type IpcLockGroups struct {
	updates int
	State   bool
}

func (fs *IpcLockGroups) Update(event, value string) bool {
	fs.updates++
	switch event {
	case "lockgroups":
		fs.State = value == "1"
	default:
		panic("wrong event: " + event)
	}

	return fs.updates == 1
}

func (fs IpcLockGroups) Target() (*bool, error) {
	return &fs.State, nil
}
