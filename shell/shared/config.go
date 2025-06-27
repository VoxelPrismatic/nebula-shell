package shared

import (
	"fmt"
	"nebula-shell/shell/shared/sakura"
)

var (
	Fonts = fontConfig{
		Sans: fontEntry{
			Name:  "Ubuntu Nerd Font",
			Scale: 1,
		},
		Mono: fontEntry{
			Name:  "GoMono Nerd Font",
			Scale: 0.9,
		},
	}
	FontScale = fontScale{
		Header: 16,
		Text:   12,
	}
	Theme = sakura.MapSwatch(sakura.Sakura.Parse(), func(c uint) string {
		return fmt.Sprintf("#%06X", c)
	})
	Radius = 12
	Width  = 96
)

type fontEntry struct {
	Name  string
	Scale float64
}
type fontConfig struct {
	Sans fontEntry
	Mono fontEntry
}
type fontScale struct {
	Header float64
	Text   float64
}
