package main

import (
	"fmt"
)

func getType(x interface{}) {
	switch t := x.(type) {
	case int:
		fmt.Println("Integer: ", t)
	case string:
		fmt.Println("String: ", t)
	case bool:
		fmt.Println("Bool: ", t)
	case chan interface{}:
		fmt.Println("Chan: ", t)
	}
}

func main() {
	a, b, c, d := 1, "Hello", true, make(chan interface{})
	getType(a)
	getType(b)
	getType(c)
	getType(d)

}
