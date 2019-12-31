package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
    var value1 string = os.Args[1]
    s := strings.Split(value1, ",")
    sum := 0
    for _, num := range s {
    	temp, _ := strconv.Atoi(num)
    	sum += temp
	}
	fmt.Println(sum)
}
