package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
    var value1 string = os.Args[1]

    s := strings.Split(value1, ",")

    var bigones []int
	bigones = append(bigones, 0)

    for i := 1; i < len(s); i++ {
			if len(s[i]) > len(s[bigones[0]]) {
				bigones = bigones[:0]
				bigones = append(bigones, i)
			}
			if len(s[i]) == len(s[bigones[0]]) {
				bigones = append(bigones, i)
			}
	}

	for i, indx := range bigones {
		if i == 0 {
			fmt.Printf("%s", strings.ToLower(s[indx]))
		} else {
			fmt.Printf(",%s", strings.ToLower(s[indx]))
		}
	}
}