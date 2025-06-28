package hyprctl

type HyprCursorPos struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func CursorPos() (*HyprCursorPos, error) {
	return Call[HyprCursorPos]("cursorpos")
}

func (p *HyprCursorPos) MoveRel(x, y int) error {
	_, err := Dispatch("movecursor", p.X+x, p.Y+y)
	if err != nil {
		return err
	}
	p.X += x
	p.Y += y
	return nil
}

func (p *HyprCursorPos) MoveAbs(x, y int) error {
	_, err := Dispatch("movecursor", x, y)
	if err != nil {
		return err
	}
	p.X = x
	p.Y = y
	return nil
}
