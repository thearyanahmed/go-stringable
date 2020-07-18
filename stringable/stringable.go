package stringable

import (
	"strings"
	"regexp"
)

type Banana string

var (
	matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")
)

func (s *Str) SnakeCase() string {
	snake := matchFirstCap.ReplaceAllString(s, "${1}_${2}")
    snake  = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
    return strings.ToLower(snake)
}