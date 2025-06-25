package main

import "nebula-shell/svc/hypripc"

func main() {
	instance := hypripc.IpcListener{}
	err := instance.Connect()
	if err != nil {
		panic(err)
	}
}
