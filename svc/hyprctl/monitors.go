package hyprctl

import (
	"fmt"

	"github.com/mappu/miqt/qt6"
)

type HyprMonitorList []HyprMonitor

type HyprMonitorName string

type HyprMonitorRef struct {
	Id   int             `json:"id"`
	Name HyprMonitorName `json:"name"`
}

type HyprMonitor struct {
	HyprMonitorRef
	Description      string           `json:"description"`
	Make             string           `json:"make"`
	Model            string           `json:"model"`
	Serial           string           `json:"serial"`
	Width            int              `json:"width"`
	Height           int              `json:"height"`
	RefreshRate      float64          `json:"refreshRate"`
	X                int              `json:"x"`
	Y                int              `json:"y"`
	ActiveWorkspace  HyprWorkspaceRef `json:"activeWorkspace"`
	SpecialWorkspace HyprWorkspaceRef `json:"specialWorkspace"`
	Reserved         []int            `json:"reserved"`
	Scale            float64          `json:"scale"`
	Transform        int              `json:"transform"`
	Focused          bool             `json:"focused"`
	DpmsStatus       bool             `json:"dpmsStatus"`
	Vrr              bool             `json:"vrr"`
	Solitary         string           `json:"solitary"`
	ActivelyTearing  bool             `json:"activelyTearing"`
	DirectScanoutTo  string           `json:"directScanoutTo"`
	Disabled         bool             `json:"disabled"`
	CurrentFormat    string           `json:"currentFormat"`
	MirrorOf         HyprMonitorName  `json:"mirrorOf"`
	AvailableModes   []string         `json:"availableModes"`
}

func Monitors() (*HyprMonitorList, error) {
	return Call[HyprMonitorList]("monitors")
}

func Monitor(name string) (*HyprMonitor, error) {
	return HyprMonitorName(name).Target()
}

func (mon HyprMonitor) ToQRect() *qt6.QRect {
	return qt6.NewQRect4(mon.X, mon.Y, mon.Width, mon.Height)
}

func (mon HyprMonitorRef) Target() (*HyprMonitor, error) {
	return mon.Name.Target()
}

func (mon HyprMonitorName) Target() (*HyprMonitor, error) {
	mons, err := Monitors()
	if err != nil {
		return nil, err
	}
	for _, m := range *mons {
		if mon == m.Name {
			return &m, nil
		}
	}
	return nil, fmt.Errorf("monitor not found")
}

func (mons *HyprMonitorList) Find(name string) *HyprMonitor {
	for _, mon := range *mons {
		if string(mon.Name) == string(name) {
			return &mon
		}
	}
	return nil
}

func (mon *HyprMonitorRef) ToQScreen() *qt6.QScreen {
	for _, screen := range qt6.QGuiApplication_Screens() {
		if screen.Name() == string(mon.Name) {
			return screen
		}
	}
	return nil
}

func (mon HyprMonitorName) Ref() *HyprMonitorRef {
	return &HyprMonitorRef{Name: mon}
}
