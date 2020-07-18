package main

import (
	"github.com/thearyanahmed/go-stringable/stringable"
)

func main() {
	x := stringable.ToStringable("The quick brown fox jumps over the lazy dog.")

	//fmt.Println(x.ToLower().SnakeCase().Get())
	p := map[string]string{
		"quick":"fast",
		"over": "above",
	}

	x.Slugify()
}

