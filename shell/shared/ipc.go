package shared

import "nebula-shell/svc/hypripc"

var instance *hypripc.IpcListener

func Ipc() *hypripc.IpcListener {
	if instance == nil {
		instance = &hypripc.IpcListener{}
		go instance.Connect()
	}
	return instance
}
