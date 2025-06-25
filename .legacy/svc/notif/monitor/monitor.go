package kdebus

import (
	"fmt"

	"github.com/godbus/dbus/v5"
)

type Monitor struct {
	Connection *dbus.Conn
	Listeners  []*Listener
}

type Listener struct {
	MatchHeader map[dbus.HeaderField]dbus.Variant
	MatchType   dbus.Type
	Callback    func(*dbus.Message)
}

func (mon *Monitor) AddListener(lis *Listener) {
	mon.Listeners = append(mon.Listeners, lis)
}

func (mon *Monitor) Connect() error {
	conn, err := dbus.ConnectSessionBus()
	if err != nil {
		return err
	}

	call := conn.BusObject().Call(
		"org.freedesktop.DBus.Monitoring.BecomeMonitor",
		0,
		[]string{"eavesdrop=true"},
		uint32(0),
	)
	if call.Err != nil {
		return fmt.Errorf("could not become monitor: %v", call.Err)
	}

	mon.Connection = conn
	return nil
}

func (mon *Monitor) Begin() {
	if mon.Connection == nil {
		err := mon.Connect()
		if err != nil {
			panic(err)
		}
	}

	msgs := make(chan *dbus.Message)
	mon.Connection.Eavesdrop(msgs)

	for msg := range msgs {
		fmt.Println(msg.String())
		for _, listener := range mon.Listeners {
			if listener.MatchType != 0 && listener.MatchType != msg.Type {
				fmt.Printf("- Did not match event type: %d != %d\n", listener.MatchType, msg.Type)
				goto nextListener
			}
			for f, val := range listener.MatchHeader {
				field := msg.Headers[f]
				if field.Signature() != val.Signature() {
					fmt.Printf("- Signature for %d did not match: %v != %v\n", f, val.Signature(), field.Signature())
					goto nextListener
				} else if field.Value() != val.Value() {
					fmt.Printf("- Value for %d did not match: %v != %v\n", f, val.Value(), field.Value())
					goto nextListener
				}
			}

			fmt.Println("- Reached callback!")
			listener.Callback(msg)

		nextListener:
			fmt.Println()
		}
	}
}
