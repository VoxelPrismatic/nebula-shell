package tiles

import "os/user"

var DesktopCache = desktopCache{}

type desktopCache struct {
}

func (c *desktopCache) Refresh() {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}
	u.HomeDir
}
