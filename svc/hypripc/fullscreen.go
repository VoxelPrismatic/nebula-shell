package hypripc

type IpcFullscreen struct {
	updates int
	State   bool
}

func (fs *IpcFullscreen) Update(event, value string) bool {
	fs.updates++
	switch event {
	case "fullscreen":
		fs.State = value == "1"
	default:
		panic("wrong event: " + event)
	}

	return fs.updates == 1
}

func (fs IpcFullscreen) Target() (*bool, error) {
	return &fs.State, nil
}
