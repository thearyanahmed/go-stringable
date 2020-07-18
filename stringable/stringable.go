package stringable

import (
	"strings"
	"unicode/utf8"
)

type Str struct {
	Value string
}

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

func (s *Str) WordCount() int {
	words := strings.Fields(s.Value)
	count := 0
	for _ = range words {
		count++
	}
	return count
}

func (s *Str) LetterCount() int {
	return utf8.RuneCountInString(s.Value)
}

func (s *Str) UCWords() *Str {
	s.Value = strings.Title(s.Value)
	return s
}

func (s *Str) After(after string) *Str {
	position := strings.LastIndex(s.Value,after)

	if position == -1 {
		return s
	}
	adjustedPosition := position + len(after)

	s.Value = s.Value[adjustedPosition:]

	return s
}

func (s *Str) Before(before string) *Str {
	position := strings.Index(s.Value, before)

	if position == -1 {
		return s
	}
	s.Value = s.Value[0:position]

	return s
}

// Get substring between two strings.
func (s *Str) Between(from string, to string) *Str {
	positionFirst := strings.Index(s.Value, from)

	if positionFirst == -1 {
		return s
	}
	posLast := strings.Index(s.Value, to)

	if posLast == -1 {
		return s
	}
	positionFirstAdjusted := positionFirst + len(from)

	if positionFirstAdjusted >= posLast {
		return s
	}
	s.Value = s.Value[positionFirstAdjusted:posLast]

	return s
}

func (s *Str) Append(append string) *Str {
	s.Value = s.Value + append
	return s
}

func (s *Str) Prepend(prepend string) * Str {
	s.Value = prepend + s.Value
	return s
}

func (s *Str) Contains(substr string) bool {
	return strings.Contains(s.Value,substr)
}

func (s *Str) ContainsAll(substrs []string) bool {

	for _, substr := range substrs {
		if strings.Contains(s.Value,substr) == false {
			return false
		}
	}

	return true
}

func (s *Str) Split(niddle string) []string {
	return strings.Split(s.Value,niddle)
}

func (s *Str) ReplaceAll(find,replaceWith string) *Str {
	s.Value = strings.ReplaceAll(s.Value,find,replaceWith)
	return s
}

func (s *Str) Replace(find, replaceWith string) *Str {
	s.Value = strings.Replace(s.Value,find,replaceWith,-1)
	return s
}

func (s *Str) ParseWith(data map[string]string) *Str {
	for find, replace := range data {
		s.Value = strings.Replace(s.Value,find,replace,-1)
	}
	return s
}

func (s *Str) TakeFirst(limit int) *Str {
	s.Value = s.Value[limit:]
	return s
}

func (s *Str) TakeLast(limit int) *Str {
	s.Value = s.Value[:limit]
	return s
}

func (s *Str) Slugify() *Str {
	s.Value = strings.ReplaceAll(s.Value," ","-")
	return s
}

func (s *Str) Substr(from,till int) *Str {
	s.Value = s.Value[from:till]
	return s
}

func (s *Str) Trim(trimmer string) *Str {
	s.Value = strings.Trim(s.Value,trimmer)
	return s
}

func (s *Str) Rtrim(cutset string) *Str {
	s.Value = strings.TrimLeft(s.Value,cutset)
	return s
}

func (s *Str) Ltrim(cutset string) *Str {
	s.Value = strings.TrimLeft(s.Value,cutset)
	return s
}

func (s *Str) ContainsRune(r rune) bool {
	return strings.ContainsRune(s.Value,r)
}

func (s *Str) TrimSuffix(suffix string) *Str {
	s.Value = strings.TrimSuffix(s.Value,suffix)
	return s
}

func (s *Str) LastIndex(substr string) int {
	return strings.LastIndex(s.Value,substr)
}

func (s *Str) HasSuffix(suffix string) bool {
	return strings.HasSuffix(s.Value,suffix)
}

func (s *Str) HasPrefix(suffix string) bool {
	return strings.HasPrefix(s.Value,suffix)
}

func (s *Str) ContainsAny(substr string) bool {
	return strings.ContainsAny(s.Value,substr)
}





