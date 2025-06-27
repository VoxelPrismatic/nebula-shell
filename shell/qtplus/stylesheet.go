package qtplus

import (
	"cmp"
	"fmt"
	"slices"
	"strings"
)

type Styleable interface {
	SetStyleSheet(string)
}

type Stylesheet struct {
	targets []Styleable
	styles  map[string]string
}

func (s *Stylesheet) AddTarget(target Styleable) {
	if !slices.Contains(s.targets, target) {
		s.targets = append(s.targets, target)
	}
}

func (s *Stylesheet) DropTarget(target Styleable) {
	idx := slices.Index(s.targets, target)
	if idx != -1 {
		s.targets = slices.Delete(s.targets, idx, idx+1)
	}
}

func (s *Stylesheet) Set(prop string, value any) {
	if s.styles == nil {
		s.styles = map[string]string{}
	}

	valueStr := fmt.Sprint(value)

	if valueStr == "" {
		delete(s.styles, prop)
	} else {
		s.styles[prop] = valueStr
	}

	lines := make([]string, len(s.styles))
	i := 0
	for key, val := range s.styles {
		lines[i] = fmt.Sprintf("%s: %s;", key, val)
		i++
	}
	sheet := strings.Join(lines, "\n")
	for _, target := range s.targets {
		target.SetStyleSheet(sheet)
	}
}

func (s *Stylesheet) Get(prop string) string {
	return s.styles[prop]
}

func (s *Stylesheet) Compare(prop string, expect any) int {
	return cmp.Compare(s.styles[prop], fmt.Sprint(expect))
}
