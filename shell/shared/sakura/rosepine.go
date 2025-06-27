package sakura

type HSLVector struct {
	H float64
	L float64
	S float64
}

type SakuraPaint[T any] struct {
	Love T // Red
	Gold T // Yellow
	Rose T // Pink
	Pine T // Darker blue
	Foam T // Light blue
	Iris T // Vibrant blue
	Tree T // Green
}

type SakuraHl[T any] struct {
	High T
	Med  T
	Low  T
}

type SakuraLayer[T any] struct {
	Base    T
	Overlay T
	Surface T
}

type SakuraText[T any] struct {
	Normal T
	Muted  T
	Subtle T
}

type SakuraPalette[T any] struct {
	Paint SakuraPaint[T]
	Hl    SakuraHl[T]
	Layer SakuraLayer[T]
	Text  SakuraText[T]
}

type SakuraSwatch[T any] struct {
	Dawn SakuraPalette[T]
	Moon SakuraPalette[T]
	Main SakuraPalette[T]
}

type VectorTheme[T any] struct {
	Base T // Background color
	Text T // Text color
}

type VectorPalette[T any] struct {
	Paint SakuraPaint[T]
	Dawn  VectorTheme[T]
	Moon  VectorTheme[T]
	Main  VectorTheme[T]
}

type DerivePalette VectorPalette[uint]

func MergeSwatch[T, R any](swatch SakuraSwatch[T], cb SakuraSwatch[func(T) R]) SakuraSwatch[R] {
	return SakuraSwatch[R]{
		Dawn: MergePalette(swatch.Dawn, cb.Dawn),
		Moon: MergePalette(swatch.Moon, cb.Moon),
		Main: MergePalette(swatch.Main, cb.Main),
	}
}

func MapSwatch[T, R any](swatch SakuraSwatch[T], cb func(T) R) SakuraSwatch[R] {
	return SakuraSwatch[R]{
		Dawn: MapPalette(swatch.Dawn, cb),
		Moon: MapPalette(swatch.Moon, cb),
		Main: MapPalette(swatch.Main, cb),
	}
}

func MergePalette[T, R any](p SakuraPalette[T], cb SakuraPalette[func(T) R]) SakuraPalette[R] {
	return SakuraPalette[R]{
		Paint: MergePaint(p.Paint, cb.Paint),
		Hl:    MergeHl(p.Hl, cb.Hl),
		Layer: MergeLayer(p.Layer, cb.Layer),
		Text:  MergeText(p.Text, cb.Text),
	}
}
func MapPalette[T, R any](p SakuraPalette[T], cb func(T) R) SakuraPalette[R] {
	return SakuraPalette[R]{
		Paint: MapPaint(p.Paint, cb),
		Hl:    MapHl(p.Hl, cb),
		Layer: MapLayer(p.Layer, cb),
		Text:  MapText(p.Text, cb),
	}
}

func MergePaint[T, R any](p SakuraPaint[T], cb SakuraPaint[func(T) R]) SakuraPaint[R] {
	return SakuraPaint[R]{
		Love: cb.Love(p.Love),
		Rose: cb.Rose(p.Rose),
		Gold: cb.Gold(p.Gold),
		Iris: cb.Iris(p.Iris),
		Tree: cb.Tree(p.Tree),
		Foam: cb.Foam(p.Foam),
		Pine: cb.Pine(p.Pine),
	}
}
func MapPaint[T, R any](p SakuraPaint[T], cb func(T) R) SakuraPaint[R] {
	return SakuraPaint[R]{
		Love: cb(p.Love),
		Rose: cb(p.Rose),
		Gold: cb(p.Gold),
		Iris: cb(p.Iris),
		Tree: cb(p.Tree),
		Foam: cb(p.Foam),
		Pine: cb(p.Pine),
	}
}

func MergeHl[T, R any](p SakuraHl[T], cb SakuraHl[func(T) R]) SakuraHl[R] {
	return SakuraHl[R]{
		High: cb.High(p.High),
		Med:  cb.Med(p.Med),
		Low:  cb.Low(p.Low),
	}
}

func MapHl[T, R any](p SakuraHl[T], cb func(T) R) SakuraHl[R] {
	return SakuraHl[R]{
		High: cb(p.High),
		Med:  cb(p.Med),
		Low:  cb(p.Low),
	}
}

func DeriveHl[T, R any](p T, cb SakuraHl[func(T) R]) SakuraHl[R] {
	return SakuraHl[R]{
		High: cb.High(p),
		Med:  cb.Med(p),
		Low:  cb.Low(p),
	}
}

func MergeLayer[T, R any](p SakuraLayer[T], cb SakuraLayer[func(T) R]) SakuraLayer[R] {
	return SakuraLayer[R]{
		Base:    cb.Base(p.Base),
		Overlay: cb.Overlay(p.Overlay),
		Surface: cb.Surface(p.Surface),
	}
}

func MapLayer[T, R any](p SakuraLayer[T], cb func(T) R) SakuraLayer[R] {
	return SakuraLayer[R]{
		Base:    cb(p.Base),
		Overlay: cb(p.Overlay),
		Surface: cb(p.Surface),
	}
}

func DeriveLayer[T, R any](p T, cb SakuraLayer[func(T) R]) SakuraLayer[R] {
	return SakuraLayer[R]{
		Base:    cb.Base(p),
		Overlay: cb.Overlay(p),
		Surface: cb.Surface(p),
	}
}

func MergeText[T, R any](p SakuraText[T], cb SakuraText[func(T) R]) SakuraText[R] {
	return SakuraText[R]{
		Normal: cb.Normal(p.Normal),
		Muted:  cb.Muted(p.Muted),
		Subtle: cb.Subtle(p.Subtle),
	}
}
func MapText[T, R any](p SakuraText[T], cb func(T) R) SakuraText[R] {
	return SakuraText[R]{
		Normal: cb(p.Normal),
		Muted:  cb(p.Muted),
		Subtle: cb(p.Subtle),
	}
}

func DeriveText[T, R any](p T, cb SakuraText[func(T) R]) SakuraText[R] {
	return SakuraText[R]{
		Normal: cb.Normal(p),
		Muted:  cb.Muted(p),
		Subtle: cb.Subtle(p),
	}
}
