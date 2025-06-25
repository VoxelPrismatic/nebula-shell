package hyprctl

type HyprDevice struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

type HyprDevMouse struct {
	HyprDevice
	DefaultSpeed float64 `json:"defaultSpeed"`
}

type HyprDevKeyboard struct {
	HyprDevice
	Rules        string `json:"rules"`
	Model        string `json:"model"`
	Layout       string `json:"layout"`
	Variant      string `json:"variant"`
	Options      string `json:"options"`
	ActiveKeymap string `json:"active_keymap"`
	CapsLock     bool   `json:"capsLock"`
	NumLock      bool   `json:"numLock"`
	Main         bool   `json:"main"`
}

type HyprDevTablet struct {
	HyprDevice
	Type      string         `json:"type"`
	BelongsTo *HyprDevTablet `json:"belongsTo"`
}

type HyprDevTouch struct {
	HyprDevice
}

type HyprDevSwitch struct {
	HyprDevice
}

type HyprDeviceList struct {
	Mice      []HyprDevMouse    `json:"mice"`
	Keyboards []HyprDevKeyboard `json:"keyboards"`
	Tablets   []HyprDevTablet   `json:"tablets"`
	Touch     []HyprDevTouch    `json:"touch"`
	Switches  []HyprDevSwitch   `json:"switches"`
}

func Devices() (*HyprDeviceList, error) {
	return Call[HyprDeviceList]("devices")
}
