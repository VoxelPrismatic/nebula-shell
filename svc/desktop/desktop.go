package desktop

import (
	"fmt"
	"strings"

	"gopkg.in/ini.v1"
)

type DesktopFile struct {
	Type                 string
	Version              string
	Name                 string
	GenericName          string
	NoDisplay            bool
	Comment              string
	Icon                 string
	Hidden               bool
	OnlyShowIn           []string `delim:";"`
	NotShowIn            []string `delim:";"`
	DBusActivatable      bool
	TryExec              string
	Exec                 string
	Path                 string
	Terminal             bool
	MimeType             []string `delim:";"`
	Categories           []string `delim:";"`
	Implements           []string `delim:";"`
	Keywords             []string `delim:";"`
	StartupNotify        bool
	StartupWMClass       string
	URL                  string
	PrefersNonDefaultGPU bool
	SingleMainWindow     bool
}

type DesktopFilePlus struct {
	DesktopFile
	Actions   []*DesktopAction
	FilePath_ string
}

// var localized = []string{"name", "genericname", "comment", "keywords"}

type DesktopAction struct {
	target *DesktopFilePlus
	Name   string
	Exec   string
	Icon   string
}

func FromFile(path string) (*DesktopFilePlus, error) {
	cfg, err := ini.Load(path)
	if err != nil {
		return nil, err
	}
	ret := &DesktopFilePlus{FilePath_: path}

	isDesktop := false
	for _, sec := range cfg.Sections() {
		if strings.ToLower(sec.Name()) == "desktop entry" {
			isDesktop = true
			if err := sec.MapTo(&ret.DesktopFile); err != nil {
				panic(err)
			}
		} else {
			action := &DesktopAction{target: ret}
			if err := sec.MapTo(action); err != nil {
				panic(err)
			}
			ret.Actions = append(ret.Actions, action)
		}
	}

	if !isDesktop {
		return ret, fmt.Errorf("not a fd.desktop file")
	}
	return ret, nil
}
