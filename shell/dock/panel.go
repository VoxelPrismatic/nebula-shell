package dock

import (
	"nebula-shell/svc/layershell"

	"github.com/mappu/miqt/qt6"
)

func NewDock(screen *qt6.QScreen) {
	window := qt6.NewQWindow()
	wlr := layershell.MakeWindow(window)
	wlr.SetWlrScope("net.voxelprismatic.nebula")
	wlr.SetWlrAnchors(layershell.AnchorBottom | layershell.AnchorRight | layershell.AnchorTop)

	rect := screen.AvailableGeometry()
	w := 96
	h := rect.Height()
	wlr.SetWlrExclusionZone(int32(w))
	window.SetGeometry(rect.X()+rect.Width()-w, rect.Y()+rect.Height()-h, w, h)
	window.Show()
}
