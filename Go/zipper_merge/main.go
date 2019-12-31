package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var value1 string = os.Args[1]
	var value2 string = os.Args[2]

	s := strings.Split(value1, ",")
	t := strings.Split(value2, ",")

	var zipped []int

	for i, _ := range s {
		one, _ := strconv.Atoi(s[i])
		two, _ := strconv.Atoi(t[i])
		zipped = append(zipped, one)
		zipped = append(zipped, two)
	}

	for i, num := range zipped {
		if i == 0 {
			fmt.Printf("%d", num)
		} else {
			fmt.Printf(",%d", num)
		}
	}
}