package main

import "fmt"

func main() {
	var input string
	fmt.Println("Input string: ")
	if _, err := fmt.Scan(&input); err != nil {
		fmt.Println(err.Error())
		return
	}
	runes := []rune(input)
	reversed := ""
	for i := len(runes) - 1; i > -1; i-- {
		reversed += string(runes[i])
	}
	fmt.Println(reversed)
}
