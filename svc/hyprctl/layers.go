package hyprctl

import "strconv"

type HyprLayerRef struct {
	Namespace string `json:"namespace"`
}

type HyprLayer struct {
	HyprLayerRef
	Address string `json:"address"`
	X       int    `json:"x"`
	Y       int    `json:"y"`
	Width   int    `json:"w"`
	Height  int    `json:"h"`
	Pid     int    `json:"pid"`
}

type HyprLayerTarget struct {
	Level   int
	Monitor *HyprMonitorRef
	Layer   *HyprLayer
}

type _HyprLayerCall map[string]struct {
	Levels map[string][]HyprLayer `json:"levels"`
}

type HyprLayerCall map[HyprMonitorRef][][]HyprLayer

func Layers() (*HyprLayerCall, error) {
	obj, err := Call[_HyprLayerCall]("layers")
	if err != nil {
		return nil, err
	}

	ret := HyprLayerCall{}
	for mon, levels := range *obj {
		key := *HyprMonitorName(mon).Ref()
		val := make([][]HyprLayer, len(levels.Levels))
		for idx, layers := range levels.Levels {
			i, err := strconv.Atoi(idx)
			if err != nil {
				return &ret, err
			}
			for i >= len(val)-1 {
				val = append(val, nil)
			}
			val[i] = layers
		}
		ret[key] = val
	}
	return &ret, nil
}

func (lay HyprLayerRef) Target() (*[]HyprLayerTarget, error) {
	obj, err := Layers()
	if err != nil {
		return nil, err
	}
	ret := []HyprLayerTarget{}
	for mon, levels := range *obj {
		for level, layers := range levels {
			for _, layer := range layers {
				if lay.Namespace == layer.Namespace {
					ret = append(ret, HyprLayerTarget{
						Level:   level,
						Layer:   &layer,
						Monitor: &mon,
					})
				}
			}
		}
	}
	return &ret, nil

}
