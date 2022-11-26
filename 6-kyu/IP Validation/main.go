package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var ip string
	_, _ = fmt.Scanln(&ip)
	var flag bool
	flag = Is_valid_ip(ip)
	if flag == true {
		fmt.Println("Should test that " + "ip" + " is correct")
	} else {
		fmt.Println("Should test that " + "ip" + " is uncorrect")
	}
}

func Is_valid_ip(ip string) bool {
	var str []string
	str = strings.Split(ip, ".")
	if len(str) != 4 {
		return false
	}
	for _, i := range str {
		num, err := strconv.Atoi(i)
		if err != nil {
			return false
		}
		if num <= -1 || num >= 256 {
			return false
		}
		s := strconv.Itoa(num)
		if len(i) != len(s) {
			return false
		}
	}
	return true
}
