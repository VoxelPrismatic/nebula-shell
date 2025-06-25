package main

import "nebula-shell/svc/hypripc"

func main() {
	err := hypripc.Connect()
	if err != nil {
		panic(err)
	}
}
