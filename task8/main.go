package main

import (
	"fmt"
)

func main() {
	var num, pos, sym int64
	fmt.Println("Input num: ")
	fmt.Scan(&num)
	fmt.Println("Input pos: ")
	fmt.Scan(&pos)
	fmt.Println("Input sym: ")
	fmt.Scan(&sym)
	mask := int64(1)<<pos - 1
	if sym == 0 {
		num &^= mask
	} else {
		num |= mask
	}
	fmt.Println(num)
}
