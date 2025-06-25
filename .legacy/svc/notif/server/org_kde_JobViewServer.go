package server

import (
	"fmt"
	kdebus "kdebus/monitor"
	"reflect"

	"github.com/godbus/dbus/v5"
)

// type KdeJob struct {
// Source string
// Destination string
// DestUrl string
// Immediate bool
// Title string
// TotalBytes uint64
// TotalFiles uint64
// ElapsedTime int64
// ProcessedBytes uint64
// ProcessedFiles uint64
// Speed uint64
// Percent uint32
// }
//
// var Jobs = map[string]KdeJob{}
// var WaitingForId = map[uint32]KdeJob{}
//
// var UpdateTask = kdebus.Listener{
// MatchHeader: map[dbus.HeaderField]dbus.Variant{
// dbus.FieldDestination: dbus.MakeVariant("org.kde.JobViewServer"),
// dbus.FieldInterface: dbus.MakeVariant("org.kde.JobViewV3"),
// dbus.FieldMember: dbus.MakeVariant("update"),
// },
// MatchType: dbus.TypeMethodCall,
// Callback: func(m *dbus.Message) {
//
// },
// }

var CreateTask = kdebus.Listener{
	MatchHeader: map[dbus.HeaderField]dbus.Variant{
		dbus.FieldDestination: dbus.MakeVariant("org.kde.JobViewServer"),
		dbus.FieldInterface:   dbus.MakeVariant("org.kde.JobViewServerV2"),
		dbus.FieldMember:      dbus.MakeVariant("requestView"),
	},
	MatchType: dbus.TypeMethodCall,
	Callback: func(m *dbus.Message) {
		var content map[string]dbus.Variant
		for _, part := range m.Body {
			if reflect.ValueOf(part).Kind() != reflect.Map {
				continue
			}

			content = part.(map[string]dbus.Variant)
			break
		}

		for key, val := range content {
			fmt.Printf("    \x1b[94;1m%s\x1b[0m: %s\n", key, val.Value())
		}
	},
}
