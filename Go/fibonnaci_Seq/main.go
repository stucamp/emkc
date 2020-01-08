package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

var PHI float64 = (math.Sqrt(5) + 1) / 2
var phi float64 = (-math.Sqrt(5) + 1) / 2

func main() {
	value1, _ := strconv.Atoi(os.Args[1])
	value2, _ := strconv.Atoi(os.Args[2])

	results := findNext(value1, value2)

	for i, fib := range results {
		if i == 0 {
			fmt.Printf("%d", fib)
		} else {
			fmt.Printf(",%d", fib)
		}
	}
}

func findNext(start int, n int) []int {
	beg := findIndex(start)
	var results []int
	for i := beg + 1; i <= n+beg; i++ {
		results = append(results, fib(i))
	}
	return results
}

func findIndex(n int) int {
	eps := math.Pow(10, -10)
	one := float64(n)*math.Sqrt(5) + eps
	almost := math.Log(one) / math.Log(PHI)
	return int(math.Round(almost))
}

func fib(n int) int {
	almost := (math.Pow(PHI, float64(n)) - math.Pow(phi, float64(n))) / math.Sqrt(5)
	return int(math.Round(almost))
}
