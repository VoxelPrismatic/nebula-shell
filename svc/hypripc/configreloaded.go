package hypripc

type IpcConfigReloaded struct {
	state bool
}

func (fs *IpcConfigReloaded) Update(event, value string) bool {
	return true
}

func (fs IpcConfigReloaded) Target() (*bool, error) {
	return &fs.state, nil
}
