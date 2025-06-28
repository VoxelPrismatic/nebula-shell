package hyprctl

import "fmt"

type HyprWorkspaceRef struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type HyprWorkspace struct {
	HyprWorkspaceRef
	Monitor         HyprMonitorName `json:"monitor"`
	MonitorId       int             `json:"monitorID"`
	Windows         int             `json:"windows"`
	HasFullscreen   bool            `json:"hasfullscreen"`
	LastWindow      HyprWindowAddr  `json:"lastwindow"`
	LastWindowTitle string          `json:"lastwindowtitle"`
	IsPersistent    bool            `json:"ispersistent"`
}

func (ws HyprWorkspace) GetLastWindow() (*HyprWindow, error) {
	return HyprWindowRef{Address: ws.LastWindow, Title: ws.LastWindowTitle}.Target()
}

func (ws HyprWorkspace) GetMonitor() (*HyprMonitor, error) {
	return HyprMonitorRef{Id: ws.MonitorId, Name: ws.Monitor}.Target()
}

func (ws HyprWorkspaceRef) Target() (*HyprWorkspace, error) {
	wss, err := Workspaces()
	if err != nil {
		return nil, err
	}
	for _, w := range *wss {
		if w.Id == ws.Id || (ws.Id == 0 && w.Name == ws.Name) {
			return &w, nil
		}
	}
	return nil, fmt.Errorf("workspace not found")
}

func Workspaces() (*[]HyprWorkspace, error) {
	return Call[[]HyprWorkspace]("workspaces")
}

func ActiveWorkspace() (*HyprWorkspace, error) {
	return Call[HyprWorkspace]("activeworkspace")
}
