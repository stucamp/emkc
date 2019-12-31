package main

import (
	"fmt"
	"os"
)

var m = map[rune]int{
'I':1,
'V':5,
'X':10,
'L':50,
'C':100,
'D':500,
'M':1000,
}

func main() {
	var value1 string = os.Args[1]
	r := []rune(value1)
	fmt.Printf("%d\n", parseRomanNums(r))
}

func parseRomanNums(nums []rune) int {
	tot := 0
	if len(nums) == 1 {
		tot += m[nums[0]]
	} else {
		for i := len(nums) - 1; i >= 0; i-- {
			if i != 0 {
				if m[nums[i]] > m[nums[i-1]] {
					tot += m[nums[i]]
					tot -= m[nums[i-1]]
					i--
				} else {
					tot += m[nums[i]]
				}
			} else if i == 0 {
				if m[nums[0]] < m[nums[1]] {
					tot -= m[nums[0]]
				} else {
					tot += m[nums[0]]
				}
			}
		}
	}
	return tot
}