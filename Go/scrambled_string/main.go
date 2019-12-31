package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

type ByRune []rune

func (r ByRune) Len() int           { return len(r) }
func (r ByRune) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r ByRune) Less(i, j int) bool { return r[i] < r[j] }

func main() {
    var value1 string = os.Args[1]
    temp := strings.ToLower(value1)
    s := strings.Split(temp, ",")
    var one ByRune = []rune(s[0])
    sort.Sort(one)
    var two ByRune = []rune(s[1])
    sort.Sort(two)

    if isEqual(one, two) {
    	fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}

func isEqual(a []rune, b []rune) bool {
	if len(a) != len(b) {
		return false
	} else {
		for i, _ := range a {
			if a[i] != b[i] {
				return false
			}
		}
		return true
	}
}