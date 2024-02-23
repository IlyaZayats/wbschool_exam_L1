package main

import (
	"fmt"
	"math"
	"wbl1/task24/point"
)

func main() {
	a, b := point.NewPoint(5, 0), point.NewPoint(0, 5)
	distance := math.Sqrt(math.Pow(b.X-a.X, 2) + math.Pow(b.Y-a.Y, 2))
	fmt.Println(distance)
}
