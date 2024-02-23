package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "snow dog sun"
	s := strings.Split(input, " ")
	fmt.Println(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	output := strings.Join(s, " ")
	fmt.Println(output)
}
