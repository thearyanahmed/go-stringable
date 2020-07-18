package stringable

import (
	"strings"
)

type Str struct {
	Value string
}
//
//var (
//	matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
//	matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")
//)

func ToStringable(str string) *Str {
	return &Str{Value: str}
}

func (s *Str) ToLower() *Str {
	s.Value = strings.ToLower(s.Value)
	return s
}

func (s *Str) SnakeCase() *Str {
	s.Value = strings.ReplaceAll(s.Value," ","_")
	return s
}

func (s *Str) Get() string {
	return s.Value
}