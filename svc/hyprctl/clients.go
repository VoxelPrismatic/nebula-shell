package hyprctl

import (
	"fmt"
	"nebula-shell/svc/desktop"
	"os"
	"path/filepath"
	"strings"
)

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

func (win *HyprWindow) DesktopFile() *desktop.DesktopFilePlus {
	if v := win.bestByName(); v != nil {
		return v
	}
	if v := win.bestByBinary(); v != nil {
		return v
	}
	if v := win.bestByClass(); v != nil {
		return v
	}
	if v := win.bestByMatch(); v != nil {
		return v
	}
	return nil
}

func (win *HyprWindow) Binary() (string, error) {
	link, err := os.Readlink(fmt.Sprintf("/proc/%d/exe", win.Pid))
	if err != nil {
		return "", err
	}
	return filepath.Base(link), nil
}

func (win *HyprWindow) bestByName() *desktop.DesktopFilePlus {
	c := win.Class + ".desktop"
	later := map[string]*desktop.DesktopFilePlus{}
	for path, obj := range desktop.Cache.Cache {
		if !strings.Contains(path, "/applications/") {
			later[path] = obj
			continue
		}
		f := filepath.Base(path)
		if strings.EqualFold(f, c) {
			return obj
		}
	}
	for path, obj := range later {
		f := filepath.Base(path)
		if strings.EqualFold(f, c) {
			return obj
		}
	}
	return nil
}

func (win *HyprWindow) bestByBinary() *desktop.DesktopFilePlus {
	bin, _ := win.Binary()
	if bin == "" {
		return nil
	}
	for _, obj := range desktop.Cache.Cache {
		if obj.Exec == "" {
			continue
		}

		fields := strings.Fields(obj.Exec)
		if len(fields) == 0 {
			continue
		}

		if strings.EqualFold(fields[0], bin) {
			return obj
		}
	}
	return nil
}

func (win *HyprWindow) bestByClass() *desktop.DesktopFilePlus {
	for _, obj := range desktop.Cache.Cache {
		if strings.EqualFold(obj.StartupWMClass, win.InitialClass) {
			return obj
		}
	}
	return nil
}

func (win *HyprWindow) bestByMatch() *desktop.DesktopFilePlus {
	parts := strings.Split(win.Class, ".")
	c := parts[len(parts)-1]

	return desktop.Cache.Fuzzy(map[string]func(f *desktop.DesktopFilePlus) string{
		c: func(f *desktop.DesktopFilePlus) string {
			return f.Name
		},
	})
}
