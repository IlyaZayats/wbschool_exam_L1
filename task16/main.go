package main

import (
	"fmt"
	"sort"
)

func main() {
	slice := []int{1, 7, 3, 10, 4, 8, 2, 5, 9, 6}
	sort.Sort(sort.IntSlice(slice))
	fmt.Println(slice)
}
