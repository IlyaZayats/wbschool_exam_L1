package main

import "fmt"

func main() {
	seq := []string{"cat", "cat", "dog", "cat", "tree"}
	a := map[string]struct{}{}
	for _, item := range seq {
		a[item] = struct{}{}
	}
	fmt.Println(a)
}
