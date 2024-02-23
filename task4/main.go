package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func main() {
	var N int

	sig := make(chan os.Signal)

	fmt.Println("Input number of workers: ")
	if _, err := fmt.Scan(&N); err != nil {
		fmt.Println(err.Error())
		return
	}

	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	ctx := context.Background()
	childCtx, cancel := context.WithCancel(ctx)
	ch := make(chan string)

	go func() {
		sig := <-sig
		fmt.Println(fmt.Sprintf("\nExiting with signal: %v", sig))
		cancel()
		return
	}()

	for i := 0; i < N; i++ {
		go func(id int, ctx context.Context) {
			for {
				select {
				case msg := <-ch:
					{
						fmt.Println(fmt.Sprintf("Worker %v got your message: %s", id, msg))
					}
				case <-ctx.Done():
					{
						fmt.Println(fmt.Sprintf("Worker %v stopped", id))
						return
					}
				}
			}
		}(i+1, childCtx)
	}

	for {
		select {
		case <-childCtx.Done():
			{
				time.Sleep(100 * time.Millisecond)
				fmt.Println("Main is stopped")
				return
			}
		default:
			{
				msg := RandStringRunes(10)
				for i := 0; i < N; i++ {
					ch <- msg
				}
			}
		}
	}

}
