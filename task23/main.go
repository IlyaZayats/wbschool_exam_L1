package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(slice)
	index := 3
	slice = append(slice[:index], slice[index+1:]...)
	fmt.Println(slice)
}
