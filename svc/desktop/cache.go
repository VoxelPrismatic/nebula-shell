package desktop

import (
	"github.com/lithammer/fuzzysearch/fuzzy"
)

type desktopCache struct {
	Cache map[string]*DesktopFilePlus
}

var Cache = desktopCache{}

func (d *desktopCache) Clear() {
	d.Cache = map[string]*DesktopFilePlus{}
}

func (d *desktopCache) Scan() {
	if d.Cache == nil {
		d.Cache = map[string]*DesktopFilePlus{}
	}

	for _, path := range Scan() {
		if d.Cache[path] != nil {
			continue
		}

		f, err := FromFile(path)
		if err != nil {
			continue
		}
		d.Cache[path] = f
	}
}

func (d *desktopCache) find(cb func(f *DesktopFilePlus) bool) *DesktopFilePlus {
	for _, p := range d.Cache {
		if p != nil && cb(p) {
			return p
		}
	}
	return nil
}

func (d *desktopCache) Find(cbs ...func(f *DesktopFilePlus) bool) *DesktopFilePlus {
	d.Scan()
	for _, cb := range cbs {
		if v := d.find(cb); v != nil {
			return v
		}
	}
	return nil
}

func (d *desktopCache) Best(cb func(f *DesktopFilePlus) float64) (*DesktopFilePlus, float64) {
	retScore := 0.0
	retVal := &DesktopFilePlus{}
	d.Scan()
	for _, p := range d.Cache {
		if p == nil {
			continue
		}
		score := cb(p)
		if score > retScore {
			retVal = p
			retScore = score
		}
	}
	return retVal, retScore
}

func (d *desktopCache) Fuzzy(matches map[string]func(f *DesktopFilePlus) string) *DesktopFilePlus {
	d.Scan()
	prevBest := 10000000
	var best *DesktopFilePlus
	for match, key := range matches {
		targets := map[string]*DesktopFilePlus{}
		keys := []string{}
		for _, p := range d.Cache {
			if p == nil {
				continue
			}

			v := key(p)
			if _, ok := targets[v]; !ok {
				keys = append(keys, v)
				targets[v] = p
			}
		}
		ranks := fuzzy.RankFindFold(match, keys)
		if len(ranks) > 0 && ranks[0].Distance <= prevBest {
			prevBest = ranks[0].Distance
			best = targets[ranks[0].Target]
		}
	}
	return best
}
