package qtplus

import (
	"fmt"
	"slices"
	"strings"

	"github.com/mappu/miqt/qt6"
)

type Styleable interface {
	Style() *qt6.QStyle
	SetStyle(*qt6.QStyle)
	SetStyleSheet(string)
}

type Stylesheet struct {
	style   map[string]string
	targets []Styleable
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

func (s *Stylesheet) Update(prop string, value any) {
	if s.style == nil {
		s.style = map[string]string{}
	}

	valueStr := fmt.Sprint(value)

	if valueStr == "" {
		delete(s.style, prop)
	} else {
		s.style[prop] = valueStr
	}
}

func (s *Stylesheet) Apply() {
	lines := make([]string, len(s.style))
	i := 0
	for key, val := range s.style {
		lines[i] = fmt.Sprintf("%s: %s;", key, val)
		i++
	}
	sheet := strings.Join(lines, "\n")
	for _, target := range s.targets {
		target.SetStyleSheet(sheet)
	}
}

func (s *Stylesheet) Set(prop string, value any) {
	s.Update(prop, value)
	s.Apply()
}
