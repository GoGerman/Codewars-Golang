package main

import "fmt"

func main() {
	var name, grad string
	_, _ = fmt.Scanln(&name)
	grad = Greet(name)
	fmt.Println(grad)
}

func Greet(name string) string {
	if name == "Johnny" {
		return "Hello, my love!"
	}
	return "Hello, " + name + "!"
}
