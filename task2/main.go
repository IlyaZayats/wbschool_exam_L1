package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	array := []int{2, 4, 6, 8, 10}
	for _, item := range array {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			fmt.Println(num * num)
		}(item)
	}
	wg.Wait()
	//wg.Add(1)
	//go func() {
	//	for i := range ch {
	//		fmt.Println(i)
	//	}
	//	close(ch)
	//	wg.Done()
	//}()
	//wg.Wait()
}
