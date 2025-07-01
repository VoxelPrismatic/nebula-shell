package shared

import (
	"sync"

	"github.com/mappu/miqt/qt6"
	"github.com/mappu/miqt/qt6/mainthread"
)

type iconCache map[string]map[int]*qt6.QPixmap

var IconCache = iconCache{}
var iconMut sync.Mutex

func (c *iconCache) Get(name string, sz int) *qt6.QPixmap {
	iconMut.Lock()
	defer iconMut.Unlock()

	sizeCache, ok := (*c)[name]
	if !ok || sizeCache == nil {
		sizeCache = map[int]*qt6.QPixmap{}
		(*c)[name] = sizeCache
	}

	pix, ok := sizeCache[sz]
	if !ok || pix == nil {
		mainthread.Wait(func() {
			var icon *qt6.QIcon
			if qt6.QIcon_HasThemeIcon(name) {
				icon = qt6.QIcon_FromTheme(name)
			} else {
				icon = qt6.QIcon_FromTheme("folder-important")
			}
			pix = icon.Pixmap2(sz, sz)
			sizeCache[sz] = pix
		})
	}
	return pix
}
