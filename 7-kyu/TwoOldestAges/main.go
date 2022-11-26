package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, x int
	_, _ = fmt.Scan(&n)
	var sl []int
	for i := 1; i <= n; i++ {
		_, _ = fmt.Scan(&x)
		sl = append(sl, x)
	}
	var ages1 [2]int
	ages1 = TwoOldestAges(sl)
	fmt.Println(ages1)
}

func TwoOldestAges(ages []int) [2]int {
	sort.Ints(ages)
	var num1, num2 int
	num1, num2 = TwoNumbers(ages)
	return [2]int{num1, num2}
}

func TwoNumbers(ages []int) (int, int) {
	var num1, num2 int
	num1 = ages[len(ages)-2]
	num2 = ages[len(ages)-1]
	return num1, num2
}
