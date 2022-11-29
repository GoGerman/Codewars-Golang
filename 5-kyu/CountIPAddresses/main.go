package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var ip1, ip2 string
	ip1, ip2 = scan(), scan()
	var sum int
	sum = IpsBetween(ip1, ip2)
	fmt.Println(sum)
}
func scan() string {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	return in.Text()
}

func IpsBetween(start, end string) int {
	var ip1str, ip2str []string
	ip1str = strings.Split(start, ".")
	ip2str = strings.Split(end, ".")
	var ip1int, ip2int []int
	for _, i := range ip1str {
		num, _ := strconv.Atoi(i)
		ip1int = append(ip1int, num)
	}
	for _, i := range ip2str {
		num, _ := strconv.Atoi(i)
		ip2int = append(ip2int, num)
	}
	var sum int
	len2 := len(ip2int)
	for j := 0; j != len2; j++ {
		if ip1int[j] != ip2int[j] {
			switch {
			case j == 0:
				if ip1int[j] > ip2int[j] {
					sum -= (ip1int[j] - ip2int[j]) * 256 * 256 * 256
				} else {
					sum += (ip2int[j] - ip1int[j]) * 256 * 256 * 256
				}
			case j == 1:
				if ip1int[j] > ip2int[j] {
					sum -= (ip1int[j] - ip2int[j]) * 256 * 256
				} else {
					sum += (ip2int[j] - ip1int[j]) * 256 * 256
				}
			case j == 2:
				if ip1int[j] > ip2int[j] {
					sum -= (ip1int[j] - ip2int[j]) * 256
				} else {
					sum += (ip2int[j] - ip1int[j]) * 256
				}
			case j == 3:
				if ip1int[j] > ip2int[j] {
					sum -= ip1int[j] - ip2int[j]
				} else {
					sum += ip2int[j] - ip1int[j]
				}
			default:
				continue
			}
		} else {
			continue
		}
	}
	return sum
}
