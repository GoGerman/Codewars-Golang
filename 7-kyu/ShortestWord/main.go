package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var str string
	str, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	str = strings.Trim(str, "\n")
	count := FindShort(str)
	fmt.Println(count)
}

func FindShort(s string) int {
	str := strings.Split(s, " ")
	var count, count2 int = 0, 200
	for _, rg := range str {
		count = len(rg)
		if count < count2 {
			count2 = count
		}
	}
	return count2
}
