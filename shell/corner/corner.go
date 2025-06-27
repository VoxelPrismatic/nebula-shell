package corner

import (
	"github.com/mappu/miqt/qt6"
)

type CornerRadius struct {
	*qt6.QWidget

	Size  int
	Color string
}

type Corners int

const (
	CornerTopLeft Corners = 1 << iota
	CornerTopRight
	CornerBotRight
	CornerBotLeft
)

func NewCorner(size int, corners Corners) *CornerRadius {
	ret := &CornerRadius{
		QWidget: qt6.NewQWidget(nil),
		Size:    size,
	}
	sz := float64(size)

	ret.SetFixedSize2(size*2, size*2)

	sqPath := qt6.NewQPainterPath()
	sqPath.AddRect2(0, 0, sz*2, sz*2)
	ccPath := qt6.NewQPainterPath()
	ccPath.AddEllipse2(0, 0, sz*2, float64(sz*2))

	negPath := qt6.NewQPainterPath()
	if corners&CornerTopLeft == 0 {
		negPath.AddRect2(0, 0, sz, sz)
	}
	if corners&CornerTopRight == 0 {
		negPath.AddRect2(sz, 0, sz, sz)
	}
	if corners&CornerBotLeft == 0 {
		negPath.AddRect2(0, sz, sz, sz)
	}
	if corners&CornerBotRight == 0 {
		negPath.AddRect2(sz, sz, sz, sz)
	}

	invPath := sqPath.Subtracted(ccPath).Subtracted(negPath)

	ret.OnPaintEvent(func(super func(event *qt6.QPaintEvent), event *qt6.QPaintEvent) {
		brush := qt6.NewQBrush3(qt6.NewQColor6(ret.Color))
		p := qt6.NewQPainter2(ret.QPaintDevice)
		p.SetRenderHint(qt6.QPainter__Antialiasing)
		p.FillPath(invPath, brush)
		p.End()
	})
	return ret
}
