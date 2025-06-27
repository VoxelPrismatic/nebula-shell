package hyprctl

import "fmt"

type HyprMonitorList []HyprMonitor

type HyprMonitorRef struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
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
	MirrorOf         string           `json:"mirrorOf"`
	AvailableModes   []string         `json:"availableModes"`
}

func Monitors() (*HyprMonitorList, error) {
	return Call[HyprMonitorList]("monitors")
}

func Monitor(name string) (*HyprMonitor, error) {
	return HyprMonitorRef{Name: name}.Target()
}

func (mon HyprMonitorRef) Target() (*HyprMonitor, error) {
	mons, err := Monitors()
	if err != nil {
		return nil, err
	}
	for _, m := range *mons {
		if mon.Name == m.Name {
			return &m, nil
		}
	}
	return nil, fmt.Errorf("monitor not found")
}

func (mons *HyprMonitorList) Find(name string) *HyprMonitor {
	for _, mon := range *mons {
		if mon.Name == name {
			return &mon
		}
	}
	return nil
}
