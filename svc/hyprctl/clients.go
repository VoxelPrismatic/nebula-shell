package hyprctl

import "fmt"

type HyprWindowAddr string

type HyprWindowRef struct {
	Address HyprWindowAddr `json:"address"`
	Title   string         `json:"title"`
}

type HyprWindow struct {
	HyprWindowRef
	Mapped           bool             `json:"mapped"`
	Hidden           bool             `json:"hidden"`
	At               [2]int           `json:"at"`
	Size             [2]int           `json:"size"`
	Workspace        HyprWorkspaceRef `json:"workspace"`
	Floating         bool             `json:"floating"`
	Pseudo           bool             `json:"pseudo"`
	Monitor          int              `json:"monitor"`
	Class            string           `json:"class"`
	InitialClass     string           `json:"initialClass"`
	InitialTitle     string           `json:"initialTitle"`
	Pid              int              `json:"pid"`
	Xwayland         bool             `json:"xwayland"`
	Pinned           bool             `json:"pinned"`
	Fullscreen       int              `json:"fullscreen"`
	FullscreenClient int              `json:"fullscreenClient"`
	Grouped          []any            `json:"grouped"`
	Tags             []any            `json:"tags"`
	Swallowing       string           `json:"swallowing"`
	FocusHistoryId   int              `json:"focusHistoryID"`
	InhibitingIdle   bool             `json:"inhibitingIdle"`
	XdgTag           string           `json:"xdgTag"`
	XdgDescription   string           `json:"xdgDescription"`
}

func Clients() (*[]HyprWindow, error) {
	return Call[[]HyprWindow]("clients")
}

func ActiveWindow() (*HyprWindow, error) {
	return Call[HyprWindow]("activewindow")
}

func (win HyprWindowRef) Target() (*HyprWindow, error) {
	return win.Address.Target()
}

func (addr HyprWindowAddr) Target() (*HyprWindow, error) {
	clients, err := Clients()
	if err != nil {
		return nil, err
	}
	for _, c := range *clients {
		if c.Address == addr {
			return &c, nil
		}
	}
	return nil, fmt.Errorf("client not found")
}

func Client(addr string) (*HyprWindow, error) {
	return HyprWindowAddr(addr).Target()
}

func (addr HyprWindowAddr) Ref() *HyprWindowRef {
	return &HyprWindowRef{Address: addr}
}
