package hyprctl

type HyprBind struct {
	Locked         bool   `json:"locked"`
	Mouse          bool   `json:"mouse"`
	Release        bool   `json:"release"`
	Repeat         bool   `json:"repeat"`
	LongPress      bool   `json:"longPress"`
	NonConsuming   bool   `json:"non_consuming"`
	HasDescription bool   `json:"has_description"`
	ModMask        int    `json:"modmask"`
	Submap         string `json:"submap"`
	Key            string `json:"key"`
	KeyCode        int    `json:"keycode"`
	CatchAll       bool   `json:"catch_all"`
	Description    string `json:"description"`
	Dispatcher     string `json:"dispatcher"`
	Arg            string `json:"arg"`
}

func Binds() (*[]HyprBind, error) {
	return Call[[]HyprBind]("binds")
}
