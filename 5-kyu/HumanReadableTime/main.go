package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int
	_, _ = fmt.Scan(&num)
	time := HumanReadableTime(num)
	fmt.Println(time)
}

func HumanReadableTime(seconds int) string {
	var num1, num2, num3 int
	schet := seconds
	for i := schet; i > 3600; i -= 3600 {
		num1++
		schet -= 3600
	}
	for i := schet; i >= 60; i -= 60 {
		num2++
		if num2 == 60 {
			num1++
			num2 = 0
		}
		schet -= 60
	}
	num3 = schet
	str1, str2, str3 := strconv.Itoa(num1), strconv.Itoa(num2), strconv.Itoa(num3)
	str1, str2, str3 = test(num1), test(num2), test(num3)
	time := str1 + ":" + str2 + ":" + str3
	return time
}

func test(num int) string {
	str := strconv.Itoa(num)
	if num < 10 && num != 0 {
		str = "0" + str
	}
	if num == 0 {
		str = "00"
	}
	return str
}
