package layershell

/*
#cgo LDFLAGS: -L. -llayer_shell -L/usr/include/LayerShellQt -L/usr/lib64
#include "layer_shell.h"
*/
import "C"
import (
	"unsafe"

	"github.com/mappu/miqt/qt6"
)

type Layer int            // Layer represents the layer type for the layer shell.
type Anchor int           // Anchor represents the anchor positions for the layer shell.
type KbdInteractivity int // KeyboardInteractivity represents keyboard focus behavior.
type ScreenCfg int

const (
	LayerBackground Layer = iota
	LayerBottom
	LayerTop
	LayerOverlay
)

const (
	AnchorNone Anchor = 0
	AnchorTop  Anchor = 1 << (iota - 1)
	AnchorBottom
	AnchorLeft
	AnchorRight
)

const (
	KbdInteractivityNone KbdInteractivity = iota
	KbdInteractivityExclusive
	KbdInteractiveOnDemand
)

const (
	ScreenFromQWindow ScreenCfg = iota
	ScreenFromCompositor
)

func UseLayerShell() {
	C.UseLayerShell()
}

type Window struct {
	*qt6.QWindow
}

func MakeWindow(target *qt6.QWindow) *Window {
	return &Window{target}
}

func (win *Window) SetWlrAnchors(anchors Anchor) {
	C.WinSetAnchors(win.UnsafePointer(), C.int(anchors))
}

func (win *Window) WlrAnchors() Anchor {
	return Anchor(C.WinGetAnchors(win.UnsafePointer()))
}

func (win *Window) SetWlrExclusionZone(zone int32) {
	C.WinSetExclusionZone(win.UnsafePointer(), C.int32_t(zone))
}

func (win *Window) WlrExclusionZone() int32 {
	return int32(C.WinGetExclusionZone(win.UnsafePointer()))
}

func (win *Window) SetWlrExclusiveEdge(edge Anchor) {
	C.WinSetExclusiveEdge(win.UnsafePointer(), C.int(edge))
}

func (win *Window) WlrExclusiveEdge() Anchor {
	return Anchor(C.WinGetExclusiveEdge(win.UnsafePointer()))
}

func (win *Window) SetWlrMargins(margins *qt6.QMargins) {
	C.WinSetMargins(win.UnsafePointer(), margins.UnsafePointer())
}

func (win *Window) SetWlrMargins2(left, top, right, bottom int) {
	win.SetWlrMargins(qt6.NewQMargins2(left, top, right, bottom))
}

func (win *Window) WlrMargins() *qt6.QMargins {
	ptr := C.WinGetMargins(win.UnsafePointer())
	if ptr == nil {
		return nil
	}

	return qt6.UnsafeNewQMargins(ptr)
}

func (win *Window) SetWlrDesiredSize(size *qt6.QSize) {
	C.WinSetDesiredSize(win.UnsafePointer(), size.UnsafePointer())
}

func (win *Window) SetWlrDesiredSize2(w, h int) {
	win.SetWlrDesiredSize(qt6.NewQSize2(w, h))
}

func (win *Window) WlrDesiredSize() *qt6.QSize {
	ptr := C.WinGetMargins(win.UnsafePointer())
	if ptr == nil {
		return nil
	}

	return qt6.UnsafeNewQSize(ptr)
}

func (win *Window) SetWlrKbdInteractivty(kbd KbdInteractivity) {
	C.WinSetKeyboardInteractivity(win.UnsafePointer(), C.int(kbd))
}

func (win *Window) WlrKbdInteractivty() KbdInteractivity {
	return KbdInteractivity(C.WinGetKeyboardInteractivity(win.UnsafePointer()))
}

func (win *Window) SetWlrLayer(layer Layer) {
	C.WinSetLayer(win.UnsafePointer(), C.int(layer))
}

func (win *Window) WlrLayer() Layer {
	return Layer(C.WinGetLayer(win.UnsafePointer()))
}

func (win *Window) SetWlrScreenCfg(cfg ScreenCfg) {
	C.WinSetScreenConfiguration(win.UnsafePointer(), C.int(cfg))
}

func (win *Window) WlrScreenCfg() ScreenCfg {
	return ScreenCfg(C.WinGetScreenConfiguration(win.UnsafePointer()))
}

func (win *Window) SetWlrScope(scope string) {
	buf := C.CString(scope)
	defer C.FreeBuf(unsafe.Pointer(buf))
	C.WinSetScope(win.UnsafePointer(), buf)
}

func (win *Window) WlrScope() string {
	buf := C.WinGetScope(win.UnsafePointer())
	if buf == nil {
		return ""
	}
	defer C.FreeBuf(unsafe.Pointer(buf))

	return C.GoString(buf)
}

func (win *Window) SetWlrCloseOnDismissed(close_ bool) {
	C.WinSetCloseOnDismissed(win.UnsafePointer(), C.bool(close_))
}

func (win *Window) WlrCloseOnDismissed() bool {
	return bool(C.WinGetCloseOnDismissed(win.UnsafePointer()))
}
