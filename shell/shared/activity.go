package shared

import (
	"nebula-shell/shell/qtplus"

	"github.com/mappu/miqt/qt6"
	"github.com/mappu/miqt/qt6/mainthread"
)

type ActivityFlag int

const (
	EntryActive ActivityFlag = 1 << iota
	EntryEmpty
	EntryOpen
)

func (a ActivityFlag) All(flags ...ActivityFlag) bool {
	for _, flag := range flags {
		if a&flag == 0 {
			return false
		}
	}
	return true
}

func (a ActivityFlag) Any(flags ...ActivityFlag) bool {
	for _, flag := range flags {
		if a&flag != 0 {
			return true
		}
	}
	return false
}

type WithQFont interface {
	Font() *qt6.QFont
	SetFont(*qt6.QFont)
}

type WithQMouseArea interface {
	UnderMouse() bool
}

func (a ActivityFlag) SetColors(target WithQFont, area WithQMouseArea, style *qtplus.Stylesheet) {
	mainthread.Start(func() {
		f := target.Font()
		f.SetBold(a.Any(EntryActive, EntryEmpty))
		target.SetFont(f)

		hovered := area.UnderMouse()
		switch a {
		case EntryActive | EntryEmpty, EntryActive | EntryEmpty | EntryOpen:
			style.Set("color", Theme.Moon.Paint.Foam)
			if hovered {
				style.Set("background-color", Theme.Moon.Layer.Base)
			} else {
				style.Set("background-color", Theme.Dawn.Text.Normal)
			}
		case EntryActive, EntryActive | EntryOpen:
			style.Set("color", Theme.Dawn.Layer.Base)
			if hovered {
				style.Set("background-color", Theme.Moon.Layer.Base)
			} else {
				style.Set("background-color", Theme.Dawn.Text.Normal)
			}
		case EntryEmpty | EntryOpen:
			style.Set("color", Theme.Dawn.Layer.Base)
			if hovered {
				style.Set("background-color", Theme.Dawn.Paint.Love)
			} else {
				style.Set("background-color", Theme.Dawn.Paint.Rose)
			}
		case EntryEmpty:
			style.Set("color", Theme.Dawn.Paint.Love)
			if hovered {
				style.Set("background-color", Theme.Dawn.Hl.Med)
			} else {
				style.Set("background-color", Theme.Dawn.Layer.Overlay)
			}
		case EntryOpen:
			style.Set("color", Theme.Dawn.Text.Normal)
			if hovered {
				style.Set("background-color", Theme.Dawn.Hl.Med)
			} else {
				style.Set("background-color", Theme.Dawn.Hl.High)
			}
		default:
			style.Set("color", Theme.Dawn.Text.Normal)
			if hovered {
				style.Set("background-color", Theme.Dawn.Hl.Med)
			} else {
				style.Set("background-color", Theme.Dawn.Layer.Overlay)
			}
		}
	})

}
