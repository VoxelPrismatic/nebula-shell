package dock

import (
	"nebula-shell/shell/shared"
	"nebula-shell/svc/hyprctl"
	"nebula-shell/svc/layershell"

	"github.com/mappu/miqt/qt6"
)

func NewTray(monitor *hyprctl.HyprMonitorRef) {
	screen := monitor.ToQScreen()
	if screen == nil {
		return
	}

	window := qt6.NewQWidget(nil)
	window.SetContentsMargins(0, 0, 0, 0)
	defer window.Show()
	window.WinId()
	window.SetAttribute(qt6.WA_TranslucentBackground)
	wlr := layershell.MakeWindow(window.WindowHandle())
	wlr.SetWlrScope("net.voxelprismatic.nebula.tray")
	wlr.SetWlrAnchors(layershell.AnchorBottom | layershell.AnchorRight)
	wlr.SetWlrKbdInteractivty(layershell.KbdInteractivityNone)
	wlr.SetWlrLayer(layershell.LayerOverlay)

	rect := screen.AvailableGeometry()
	w := 512
	h := 384
	window.SetFixedSize2(w, h)
	window.SetGeometry(rect.X()+rect.Width()-w-shared.Width-shared.Radius*2, rect.Y()+rect.Height()-h, w, h)
}
