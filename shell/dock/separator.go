package dock

import (
	"fmt"
	"nebula-shell/shell/shared"

	"github.com/mappu/miqt/qt6"
)

func NewSeparator() *qt6.QWidget {
	box := qt6.NewQWidget(nil)
	layout := qt6.NewQVBoxLayout(box)
	bar := qt6.NewQWidget(nil)
	layout.AddWidget(bar)

	box.SetFixedSize2(shared.Grid.InnerWidth(), shared.Grid.Margin)
	box.SetContentsMargins(0, 0, 0, 0)
	layout.SetContentsMargins(0, 0, 0, 0)

	bar.SetFixedSize2(shared.Grid.InnerWidth(), 2)
	bar.SetStyleSheet(fmt.Sprintf(
		"border-radius: %dpx; background-color: %s;",
		shared.Grid.Margin,
		shared.Theme.Dawn.Layer.Overlay,
	))

	return box
}
