package dock

import (
	"nebula-shell/shell/shared"
	"nebula-shell/svc/layershell"

	"github.com/mappu/miqt/qt6"
)

func NewDock(screen *qt6.QScreen) {
	window := qt6.NewQWidget(nil)
	window.SetContentsMargins(0, 0, 0, 0)
	defer window.Show()
	window.WinId()
	window.SetAttribute(qt6.WA_TranslucentBackground)
	wlr := layershell.MakeWindow(window.WindowHandle())
	wlr.SetWlrScope("net.voxelprismatic.nebula#" + screen.Name())
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

	thing := qt6.NewQHBoxLayout(window)
	thing.SetContentsMargins(0, 0, 0, 0)
	thing.SetSpacing(0)

	radiusWidget := qt6.NewQWidget(nil)
	radiusWidget.SetContentsMargins(0, 0, 0, 0)
	radiusLayout := qt6.NewQVBoxLayout(radiusWidget)
	radiusLayout.SetContentsMargins(0, 0, 0, 0)
	// radiusWidget.SetFixedSize2(shared.Radius, h)
	// radiusWidget.SetFixedWidth(shared.Radius * 2)
	radiusWidget.SetFixedHeight(h)
	radiusLayout.AddWidget(NewCorner(shared.Radius, CornerTopRight).QWidget)
	radiusLayout.AddStretch()
	radiusLayout.AddWidget(NewCorner(shared.Radius, CornerBotRight).QWidget)
	thing.AddWidget(radiusWidget)

	contentWidget := qt6.NewQWidget(nil)
	contentWidget.SetContentsMargins(0, 0, 0, 0)
	contentWidget.SetFixedSize2(shared.Width, h)
	contentLayout := qt6.NewQVBoxLayout(contentWidget)
	contentLayout.ParentWidget().SetStyleSheet("background-color: " + shared.Theme.Dawn.Layer.Base)
	thing.AddWidget(contentWidget)
}
