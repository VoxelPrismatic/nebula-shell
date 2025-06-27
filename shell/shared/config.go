package shared

import (
	"fmt"
	"nebula-shell/shell/shared/sakura"

	"github.com/mappu/miqt/qt6"
)

var (
	Fonts = fontConfig{
		Sans: fontEntry{
			Name: "Ubuntu Nerd Font",
			Sizes: map[FontSize]int{
				FzHeader: 16,
				FzText:   12,
			},
		},
		Mono: fontEntry{
			Name: "GoMono Nerd Font",
			Sizes: map[FontSize]int{
				FzHeader: 16,
				FzText:   12,
			},
		},
	}
	Theme = sakura.MapSwatch(sakura.Sakura.Parse(), func(c uint) string {
		return fmt.Sprintf("#%06X", c)
	})
	Radius = 12
	Width  = Grid.OuterWidth()
	Grid   = GridLayout{
		Columns:  3,
		CellSize: 24,
		CellPad:  2,
		Gap:      4,
		Margin:   8,
	}
)

type fontEntry struct {
	Name  string
	Sizes map[FontSize]int
	cache map[int]*qt6.QFont
}
type fontConfig struct {
	Sans fontEntry
	Mono fontEntry
}
type GridLayout struct {
	Columns  int
	CellSize int
	CellPad  int
	Gap      int
	Margin   int
}

func (g GridLayout) CellWidth() int {
	return g.CellPad*2 + g.CellSize
}
func (g GridLayout) InnerWidth() int {
	return g.CellWidth()*g.Columns + (g.Columns-1)*g.Gap
}
func (g GridLayout) OuterWidth() int {
	return g.InnerWidth() + g.Margin*2
}
func (f *fontEntry) ToQFont(px int) *qt6.QFont {
	if f.cache == nil {
		f.cache = map[int]*qt6.QFont{}
	}
	if ret, ok := f.cache[px]; ok && ret != nil {
		return ret
	}
	ret := qt6.NewQFont2(f.Name)
	ret.SetPixelSize(px)
	f.cache[px] = ret
	return ret
}
func (f *fontEntry) Preset(fz FontSize) *qt6.QFont {
	return f.ToQFont(f.Sizes[fz])
}

type FontSize int

const (
	FzHeader = iota
	FzText
)
