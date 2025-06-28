package desktop

import (
	"nebula-shell/svc/environ"
	"path/filepath"
)

func Scan() []string {

	ret := []string{}
	dirs, err := environ.Env().GetList("XDG_DATA_DIRS")
	if err != nil {
		return ret
	}

	seen := map[string]bool{}
	for _, d := range dirs {
		seen[d] = true
	}

	for len(dirs) > 0 {
		dir := dirs[0]
		dirs = dirs[1:]

		glob, err := filepath.Glob(filepath.Join(dir, "*/*.desktop"))
		if err != nil {
			panic(err)
		}
		ret = append(ret, glob...)

	}
	return ret
}
