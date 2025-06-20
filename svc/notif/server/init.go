package server

import (
	kdebus "kdebus/monitor"
)

var bus = kdebus.Monitor{}

func Begin() {
	bus.AddListener(&CreateTask)
	bus.Begin()
}
