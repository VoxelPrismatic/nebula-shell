package main

import (
	"fmt"
	"nebula-shell/svc/desktop"
)

func main() {
	d := desktop.Cache.Fuzzy(map[string]func(f *desktop.DesktopFilePlus) string{
		"zen": func(f *desktop.DesktopFilePlus) string { return f.StartupWMClass },
	})
	fmt.Println(d)
}
