package main

import (
	"fmt"
)



func ifLeap(y int) bool {
	if (y % 4 == 0) {
		if (y % 100 == 0) {
			if (y % 400 == 0) {
				return true
			}
			return false
		}
		return true
	}
	return false
}

func main() {
	y := 1996
	fmt.Println(ifLeap(y))
	y = 1997
	fmt.Println(ifLeap(y))
	y = 1900
	fmt.Println(ifLeap(y))
	y = 2000
	fmt.Println(ifLeap(y))
}
