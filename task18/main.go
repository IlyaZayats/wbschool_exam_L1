package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func NewCounter() *Counter {
	return &Counter{m: sync.Mutex{}, data: 0}
}

type Counter struct {
	m    sync.Mutex
	data int
}

func (c *Counter) Inc() {
	c.m.Lock()
	c.data++
	c.m.Unlock()
}

func main() {
	var wg sync.WaitGroup
	counter := NewCounter()
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	ctx := context.Background()
	child, cancel := context.WithCancel(ctx)
	go func() {
		sig := <-sigChan
		fmt.Println("\nExiting with signal: ", sig)
		cancel()
	}()
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(ctx context.Context) {
			for {
				select {
				case <-ctx.Done():
					wg.Done()
					return
				default:
					counter.Inc()
				}
			}
		}(child)
	}
	wg.Wait()
	fmt.Println("Counter value: ", counter.data)
}
