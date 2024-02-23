package main

import (
	"fmt"
	"sort"
)

func main() {
	slice := []int{1, 7, 3, 10, 4, 8, 2, 5, 9, 6}
	n := 2
	sort.Sort(sort.IntSlice(slice))
	index := sort.Search(len(slice), func(index int) bool {
		return slice[index] >= n
	})
	fmt.Println(index)
}
