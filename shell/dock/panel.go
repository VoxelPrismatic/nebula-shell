package dock

import (
	"fmt"
	"nebula-shell/shell/corner"
	"nebula-shell/shell/shared"
	"nebula-shell/shell/spaces"
	"nebula-shell/shell/tiles"
	"nebula-shell/svc/hyprctl"
	"nebula-shell/svc/hypripc"
	"nebula-shell/svc/layershell"

	"github.com/mappu/miqt/qt6"
)

func NewDock(monitor *hyprctl.HyprMonitorRef) {
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
	wlr.SetWlrScope("net.voxelprismatic.nebula#" + string(monitor.Name))
	wlr.SetWlrAnchors(layershell.AnchorBottom | layershell.AnchorRight | layershell.AnchorTop)
	wlr.SetWlrKbdInteractivty(layershell.KbdInteractivityNone)
	wlr.SetWlrExclusiveEdge(layershell.AnchorRight)
	wlr.SetWlrLayer(layershell.LayerBottom)

	rect := screen.AvailableGeometry()
	w := shared.Width + shared.Radius*2
	h := rect.Height()
	window.SetFixedSize2(w, h)
	window.SetScreen(screen)
	wlr.SetWlrExclusionZone(int32(shared.Width))
	window.SetGeometry(rect.X()+rect.Width()-w, rect.Y()+rect.Height()-h, w, h)

	ret := &corner.Dock{}
	corner.Docks[monitor.Name] = ret

	thing := qt6.NewQHBoxLayout(window)
	thing.SetContentsMargins(0, 0, 0, 0)
	thing.SetSpacing(0)
	ret.Thing = thing

	radiusWidget := qt6.NewQWidget(nil)
	radiusWidget.SetContentsMargins(0, 0, 0, 0)
	radiusLayout := qt6.NewQVBoxLayout(radiusWidget)
	radiusLayout.SetContentsMargins(0, 0, 0, 0)
	radiusWidget.SetFixedHeight(h)

	c := corner.NewCorner(shared.Radius, corner.CornerTopRight)
	c.Color = shared.Theme.Dawn.Layer.Base
	ret.UpperCorner = c
	radiusLayout.AddWidget(c.QWidget)
	radiusLayout.AddStretch()
	c = corner.NewCorner(shared.Radius, corner.CornerBotRight)
	c.Color = shared.Theme.Dawn.Layer.Base
	ret.LowerCorner = c
	radiusLayout.AddWidget(c.QWidget)
	thing.AddWidget(radiusWidget)

	contentWidget := qt6.NewQWidget(nil)
	contentWidget.SetContentsMargins(0, 0, 0, 0)
	contentWidget.SetFixedSize2(shared.Width, h)
	contentLayout := qt6.NewQVBoxLayout(contentWidget)
	ret.Content = contentLayout
	contentLayout.ParentWidget().SetStyleSheet(fmt.Sprintf(
		"background-color: %s;",
		shared.Theme.Dawn.Layer.Base,
	))
	gap := shared.Grid.Margin
	contentLayout.SetContentsMargins(gap, gap, gap, gap)
	thing.AddWidget(contentWidget)

	contentLayout.AddWidget(spaces.NewGrid(monitor).Widget)
	contentLayout.AddWidget(NewSeparator())
	contentLayout.AddWidget(tiles.NewGrid(monitor).Widget)
	contentLayout.AddWidget(NewSeparator())
	contentLayout.AddWidget(tiles.NewList(monitor).Widget)
	contentLayout.AddStretch()

	shared.Ipc().EvtMonitorRemoved.Add(func(imr *hypripc.IpcMonitorRemoved) bool {
		if imr.Name == monitor.Name {
			window.Close()
		}
		return true
	})
}
