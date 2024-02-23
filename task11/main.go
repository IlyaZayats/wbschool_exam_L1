package main

import "fmt"

func main() {
	a := map[int]struct{}{1: {}, 2: {}, 3: {}, 4: {}, 5: {}, 8: {}}
	b := map[int]struct{}{2: {}, 3: {}, 6: {}, 7: {}}
	c := map[int]struct{}{}
	for key := range a {
		if _, ok := b[key]; ok == true {
			c[key] = struct{}{}
		}
	}
	fmt.Println(c)
}
