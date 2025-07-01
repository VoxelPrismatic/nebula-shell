package corner

import (
	"fmt"
	"nebula-shell/svc/hyprctl"

	"github.com/mappu/miqt/qt6"
)

var Docks = map[hyprctl.HyprMonitorName]*Dock{}

type Dock struct {
	UpperCorner *CornerRadius
	LowerCorner *CornerRadius
	Content     *qt6.QVBoxLayout
	Thing       *qt6.QHBoxLayout
	color       string
}

func (d *Dock) SetColor(color string) {
	if color == d.color {
		return
	}
	d.color = color
	d.Content.ParentWidget().SetStyleSheet(fmt.Sprintf(
		"background-color: %s;", color,
	))

	d.UpperCorner.Color = color
	d.LowerCorner.Color = color
	d.LowerCorner.Repaint()
	d.UpperCorner.Repaint()
}
