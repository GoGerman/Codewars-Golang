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
	str2 := CamelCase(str)
	fmt.Println(str2)
}

func CamelCase(s string) string {
	str := strings.Split(s, " ")
	var s1 = ""
	for _, i := range str {
		s1 += strings.Title(i)
	}
	return s1
}
