package main

import "fmt"

func main() {
	var x, n int
	_, _ = fmt.Scanf("%d", &n)
	_, _ = fmt.Scanf("%d", &x)
	massiv := CountBy(x, n)
	fmt.Println(massiv)
}

func CountBy(x, n int) []int {
	var massiv []int
	for i := n; i <= n*x; i += n {
		massiv = append(massiv, i)
	}
	return massiv
}
