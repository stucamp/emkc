package main

import (
	"fmt"
	"os"
)

func main() {
	var value1 string = os.Args[1]
	r := []rune(value1)

	var mapIt = make(map[rune]int)
	mapIt['0'] = 0
	mapIt['1'] = 0
	mapIt['2'] = 0
	mapIt['3'] = 0
	mapIt['4'] = 0
	mapIt['5'] = 0
	mapIt['6'] = 0
	mapIt['7'] = 0
	mapIt['8'] = 0
	mapIt['9'] = 0

	highest := '0'

	for _, num := range r {
		mapIt[num] += 1
		if mapIt[num] >= mapIt[highest] {
			highest = num
		}
	}

	fmt.Printf("%c\n", highest)
}

//TEST
