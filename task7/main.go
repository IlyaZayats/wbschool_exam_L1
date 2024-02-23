package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	m    sync.Mutex
	data map[string]interface{}
}

func (sf *SafeMap) Get(key string) (interface{}, bool) {
	sf.m.Lock()
	defer sf.m.Unlock()
	res, k := sf.data[key]
	return res, k
}

func (sf *SafeMap) Set(key string, val interface{}) {
	sf.m.Lock()
	defer sf.m.Unlock()
	sf.data[key] = val
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		m:    sync.Mutex{},
		data: make(map[string]interface{}),
	}
}

func main() {
	safeMap := NewSafeMap()
	var wg sync.WaitGroup
	data := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	for i := 0; i < len(data); i++ {
		wg.Add(1)
		go func(id int, data []int) {
			defer wg.Done()
			safeMap.Set(fmt.Sprintf("goroutine%v", id), data)
		}(i+1, data[i])
	}
	wg.Wait()
	fmt.Println(safeMap.data)
}
