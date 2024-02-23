package main

import (
	"fmt"
	"sync"
)

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	firstChan, secondChan := make(chan int), make(chan int)
	var wg sync.WaitGroup
	go func() {
		for x := range firstChan {
			secondChan <- x * 2
		}
	}()

	go func() {
		for x := range secondChan {
			fmt.Println(x)
			wg.Done()
		}
	}()

	for _, item := range array {
		wg.Add(1)
		firstChan <- item
	}
	wg.Wait()

}
