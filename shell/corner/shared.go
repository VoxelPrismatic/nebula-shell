package corner

import (
	"nebula-shell/shell/qtplus"
	"nebula-shell/svc/hyprctl"

	"github.com/mappu/miqt/qt6"
)

var Docks = map[hyprctl.HyprMonitorName]*Dock{}

type Dock struct {
	UpperCorner *CornerRadius
	LowerCorner *CornerRadius
	Content     *qt6.QVBoxLayout
	Style       *qtplus.Stylesheet
}

func (d *Dock) SetColor(color string) {
	if d.Style.Compare("background-color", color) == 0 {
		return
	}
	d.Style.Set("background-color", color)
	d.UpperCorner.Color = color
	d.LowerCorner.Color = color
	d.LowerCorner.Repaint()
	d.UpperCorner.Repaint()
}
