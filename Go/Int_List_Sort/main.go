package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type numbers []int

func main() {
	var value1 string = os.Args[1]

	array := makeArr(value1)
	array.sortArr()

	for i := len(array)-1 ; i >= 0; i-- {
		if i == len(array)-1 {
			fmt.Printf("%d", array[i])
		} else {
			fmt.Printf(",%d", array[i])
		}
	}
}

func makeArr(n string) numbers {
	strs := strings.Split(n, ",")
	arr := make ([]int, len(strs))
	for i := range arr {
		arr[i], _ = strconv.Atoi(strs[i])
	}
	return arr
}

func (a *numbers) sortArr() {
	sort.Ints(*a)
}
