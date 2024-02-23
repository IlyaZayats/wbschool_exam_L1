package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	stop := make(chan bool)
	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("---------------------")
				fmt.Println("Stopped with channel")
				fmt.Println("---------------------")
				return
			default:
				fmt.Println("Some magic")
			}

		}
	}()
	stop <- true

	doneC := make(chan bool)
	ctx := context.Background()
	child, cancel := context.WithCancel(ctx)
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("---------------------")
				fmt.Println("Stopped with ctx")
				fmt.Println("---------------------")
				doneC <- true
				return
			default:
				fmt.Println("Some magic")
			}

		}
	}(child)
	go func() {
		cancel()
	}()
	<-doneC
	ticker := time.NewTicker(10 * time.Microsecond)
	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("---------------------")
				fmt.Println("Stopped with ticker")
				fmt.Println("---------------------")
				doneC <- true
				return
			default:
				fmt.Println("Some magic")
			}

		}
	}()
	<-doneC
	dataChan := make(chan int)
	go func() {
		<-dataChan
		fmt.Println("---------------------")
		fmt.Println("Stopped")
		fmt.Println("---------------------")
		doneC <- true
	}()
	dataChan <- 1
	<-doneC
}
