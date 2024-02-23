package main

import (
	"fmt"
	"time"
)

func main() {
	var N int
	fmt.Println("Input working time: ")
	if _, err := fmt.Scan(&N); err != nil {
		fmt.Println(err.Error())
		return
	}
	stopChan := make(chan bool)
	ticker := time.NewTicker(time.Duration(N) * time.Second)
	dataChan := make(chan int)
	go func() {
		for {
			select {
			case data := <-dataChan:
				{
					fmt.Println(fmt.Sprintf("Received data: %v", data))
				}
			case <-ticker.C:
				{
					fmt.Println("Goroutine stopped")
					close(dataChan)
					stopChan <- true
					return
				}
			}
		}
	}()

	go func() {
		data := 0
		for {
			data++
			dataChan <- data
			time.Sleep(time.Millisecond * 500)
		}
	}()
	<-stopChan
	close(stopChan)
}
