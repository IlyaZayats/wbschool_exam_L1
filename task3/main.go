package main

import "fmt"

func main() {
	array := []int{2, 4, 6, 8, 10}
	ch := make(chan int)
	sum := 0
	for _, item := range array {
		go func(sum, num int) {
			ch <- sum + (num * num)
		}(sum, item)
		sum = <-ch
	}
	fmt.Println(sum)

}
