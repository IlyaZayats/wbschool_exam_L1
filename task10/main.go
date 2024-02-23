package main

import (
	"fmt"
	"strconv"
)

func main() {
	seq := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	res := make(map[string][]float64)
	for _, item := range seq {
		key := strconv.Itoa(10 * (int(item) / 10))
		if _, ok := res[key]; ok == true {
			res[key] = append(res[key], item)
		} else {
			res[key] = []float64{item}
		}
	}
	fmt.Println(res)
}
