package qtplus

import "github.com/mappu/miqt/qt6"

type pCacheEntry struct {
	Fore int
	Back int
}

var paletteCache = map[pCacheEntry]*qt6.QPalette{}

func Palette(fore, back *qt6.QColor) *qt6.QPalette {
	e := pCacheEntry{fore.Value(), back.Value()}
	if ret, ok := paletteCache[e]; ok {
		return ret
	}

	p := qt6.NewQPalette()
	p.SetColor(qt6.QPalette__All, qt6.QPalette__Window, back)
	p.SetColor(qt6.QPalette__All, qt6.QPalette__WindowText, fore)

	paletteCache[e] = p
	return p
}
