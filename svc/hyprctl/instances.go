package hyprctl

type HyprInstance struct {
	Instance string `json:"instance"`
	Time     int    `json:"time"`
	Pid      int    `json:"pid"`
	Socket   string `json:"wl_socket"`
}

func Instances() (*[]HyprInstance, error) {
	return Call[[]HyprInstance]("instances")
}
