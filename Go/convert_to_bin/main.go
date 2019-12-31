package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var value1 string = os.Args[1]

	s := strings.Split(value1, ",")
	t := strings.Join(s, "")
	r := []rune(t)

	for _, runish := range r {
		fmt.Printf("%08b", runish)
	}
}

