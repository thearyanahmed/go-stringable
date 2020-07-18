package main

import (
	"fmt"
	"github.com/thearyanahmed/go-stringable/stringable"
)

func main() {
	x := stringable.ToStringable("Hello world")

	fmt.Println(x.ToLower().SnakeCase().Get())
}

