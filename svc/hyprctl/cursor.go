package hyprctl

type HyprCursorPos struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func CursorPos() (*HyprCursorPos, error) {
	return Call[HyprCursorPos]("cursorpos")
}
