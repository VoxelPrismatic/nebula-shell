package hyprctl

type HyprLayout string

func Layouts() (*[]HyprLayout, error) {
	return Call[[]HyprLayout]("layouts")
}
