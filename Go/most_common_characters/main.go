package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var value1 string = os.Args[1]
	input := strings.ToLower(value1)
	s := strings.Split(input, " ")
	t := strings.Join(s, "")
	r := []rune(t)

	m := map[rune]int{}

	for _, let := range r {
		_, ok := m[let]
		if ok == true {
			m[let] += 1
		} else {
			m[let] = 1
		}
	}

	var biggest []rune
	biggest = make([]rune,0)

	for key, val := range m {
		if len(biggest) == 0 {
			biggest = append(biggest, key)
		} else if val > m[biggest[0]] {
			biggest = biggest[:0]
			biggest = append(biggest, key)
		} else if val == m[biggest[0]] {
			biggest = append(biggest,key)
		}
	}

	for i, val := range biggest {
		if i == 0 {
			fmt.Printf("%c", val)
		} else {
			fmt.Printf(",%c", val)
		}
	}

}
