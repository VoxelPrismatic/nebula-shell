package qtplus

import (
	"github.com/mappu/miqt/qt6"
)

// Credit: https://stackoverflow.com/questions/10651514/text-scrolling-marquee-in-qlabel
type Marquee struct {
	*qt6.QWidget
	text            string
	separator       string
	singleTextWidth int
	wholeTextSize   *qt6.QSize
	leftMargin      int
	scrollEnabled   bool
	scrollPaused    bool
	scrollPos       int
	alphaChannel    *qt6.QImage
	buffer          *qt6.QImage
	timer           *qt6.QTimer
	staticText      *qt6.QStaticText
}

func NewMarquee(parent *qt6.QWidget) *Marquee {
	obj := qt6.NewQWidget(parent)
	ret := &Marquee{QWidget: obj}
	ret.timer = qt6.NewQTimer()
	ret.timer.SetInterval(50)
	ret.staticText = qt6.NewQStaticText()
	ret.staticText.SetTextFormat(qt6.PlainText)
	ret.SetFixedHeight(ret.FontMetrics().Height())
	ret.leftMargin = ret.Height() / 3
	ret.SetSeparator("    •     ")

	ret.timer.OnTimeout(ret.timerTimeout)

	ret.OnResizeEvent(ret.resizeEvent)
	ret.OnPaintEvent(ret.paintEvent)

	return ret
}

func (m *Marquee) SetFont(font *qt6.QFont) {
	m.QWidget.SetFont(font)
	m.updateText()
	m.Update()
}

func (m *Marquee) Text() string {
	return m.text
}

func (m *Marquee) SetText(text string) {
	if text == m.text {
		return
	}
	m.text = text
	m.updateText()
	m.Update()
}

func (m *Marquee) Separator() string {
	return m.separator
}

func (m *Marquee) SetSeparator(sep string) {
	m.separator = sep
	m.updateText()
	m.Update()
}

func (m *Marquee) Paused() bool {
	return m.scrollPaused
}

func (m *Marquee) SetPaused(paused bool) {
	if m.scrollEnabled {
		m.scrollPaused = true
		return
	}

	m.scrollPaused = paused
	if m.scrollPaused {
		m.scrollPos = 0
		m.timer.Stop()
	} else {
		m.timer.Start2()
	}
}

func (m *Marquee) updateText() {
	m.timer.Stop()

	size := m.FontMetrics().BoundingRectWithText(m.text)
	m.singleTextWidth = size.Width()
	m.scrollEnabled = m.singleTextWidth > (m.Width() - m.leftMargin)

	if m.scrollEnabled {
		m.staticText.SetText(m.text + m.separator)
		m.timer.Start2()
	} else {
		m.staticText.SetText(m.text)
	}

	m.staticText.Prepare2(qt6.NewQTransform2(), m.Font())
	size = m.FontMetrics().BoundingRectWithText(m.staticText.Text())
	m.wholeTextSize = qt6.NewQSize2(size.Width(), size.Height())
}

func (m *Marquee) paintEvent(super func(event *qt6.QPaintEvent), event *qt6.QPaintEvent) {
	p := qt6.NewQPainter2(m.QPaintDevice)
	defer p.End()

	top := (float64(m.Height()) - float64(m.wholeTextSize.Height())) / 2
	left := float64(m.leftMargin)
	topLeft := qt6.NewQPointF3(left, top)
	if !m.scrollEnabled {
		p.DrawStaticText(topLeft, m.staticText)
		return
	}

	m.buffer.Fill(0)
	pb := qt6.NewQPainter()
	pb.Begin(m.buffer.QPaintDevice)
	if pb == nil {
		panic("nil painter")
	}
	defer pb.End()
	pb.SetPenWithPen(p.Pen())
	pb.SetFont(p.Font())

	x := min(-m.scrollPos, 0) + m.leftMargin
	topLeft.SetX(float64(x))
	w, h := m.Width(), m.Height()
	for x < w {
		pb.DrawStaticText(topLeft, m.staticText)
		x += m.wholeTextSize.Width()
		topLeft.SetX(float64(x))
	}

	pb.SetCompositionMode(qt6.QPainter__CompositionMode_DestinationIn)
	pb.SetClipRect2(w-15, 0, 15, h)
	pb.DrawImage9(0, 0, m.alphaChannel)
	pb.SetClipRect2(0, 0, 15, h)
	if m.scrollPos < 0 {
		pb.SetOpacity(max(-8.0, float64(m.scrollPos)+8) / 8)
	}

	pb.DrawImage9(0, 0, m.alphaChannel)
	p.DrawImage9(0, 0, m.buffer)
}

func (m *Marquee) resizeEvent(super func(event *qt6.QResizeEvent), event *qt6.QResizeEvent) {
	m.leftMargin = m.Height() / 3
	m.alphaChannel = qt6.NewQImage2(m.Size(), qt6.QImage__Format_ARGB32_Premultiplied)
	m.buffer = qt6.NewQImage2(m.Size(), qt6.QImage__Format_ARGB32_Premultiplied)
	w := m.Width()

	if m.Width() <= 64 {
		m.alphaChannel.Fill(0)
	} else {
		for y := range m.Height() {
			for x := 1; x < 16; x++ {
				v := qt6.QRgba64_FromRgba(0, 0, 0, byte(x<<4)).ToArgb32()
				m.alphaChannel.SetPixel(x-1, y, v)
				m.alphaChannel.SetPixel(w-x, y, v)
			}
			for x := 15; x < w-15; x++ {
				m.alphaChannel.SetPixel(x, y, 0)
			}
		}
	}

	scrollEnabled := (m.singleTextWidth > w-m.leftMargin)
	if scrollEnabled != m.scrollEnabled {
		m.updateText()
	}
}

func (m *Marquee) timerTimeout() {
	m.scrollPos = (m.scrollPos + 1) % m.wholeTextSize.Width()
	m.Update()
}
