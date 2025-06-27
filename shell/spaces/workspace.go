package spaces

import (
	"nebula-shell/svc/hyprctl"

	"github.com/mappu/miqt/qt6"
)

var SpaceNames = []string{
	"α", "β", "δ",
	"ζ", "ξ", "ϟ",
	"λ", "π", "μ",
	"τ", "ω", "ϰ",
}

type WorkspaceEntry struct {
	Target     hyprctl.HyprWorkspaceRef
	Widget     *qt6.QWidget
	Monitor    *qt6.QScreen
	Stylesheet *qt6.QStyle
	Text       *qt6.QLabel
}

func NewEntry(mon *qt6.QScreen, ref hyprctl.HyprWorkspaceRef) *WorkspaceEntry {
	workspace := &WorkspaceEntry{
		Target:     ref,
		Widget:     qt6.NewQWidget2(),
		Monitor:    mon,
		Stylesheet: qt6.NewQStyle(),
		Text:       qt6.NewQLabel2(),
	}

	w := workspace.Widget
	s := workspace.Stylesheet
	w.SetStyle(workspace.Stylesheet)
	s.SetProperty("radius", qt6.NewQVariant4(4))

	layout := qt6.NewQVBoxLayout(w)
	layout.AddWidget(workspace.Text.QWidget)
	workspace.Text.SetAlignment(qt6.AlignCenter)

	w.OnMousePressEvent(func(super func(event *qt6.QMouseEvent), event *qt6.QMouseEvent) {
		switch event.Button() {
		case qt6.RightButton:
			_, _ = hyprctl.Dispatch("movetoworkspace", workspace.Target.Id)
		case qt6.LeftButton:
			_, _ = hyprctl.Dispatch("focusworkspaceoncurrentmonitor", workspace.Target.Id)
		case qt6.MiddleButton:
			_, _ = hyprctl.Dispatch("movetoworkspacesilent", workspace.Target.Id)
		}
	})

	return workspace
}
