package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var srt string = scan()
	var bo bool = ValidParentheses(srt)
	fmt.Println(bo)
}

func scan() string {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	return in.Text()
}

func ValidParentheses(parens string) bool {
	var count int
	if parens == "" {
		return true
	}
	for _, i := range parens {
		if i == 40 {
			count++
		}
		if i == 41 {
			count--
		}
		if count < 0 {
			fmt.Println("z")
			return false
		}
	}
	if count == 0 {
		return true
	} else {
		return false
	}
}
