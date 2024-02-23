package main

import (
	"fmt"
	"math/rand"
	"time"
	"unsafe"
)

var justString string

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandStringRunes(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func Clone(str string, n int) string {
	if len(str) == 0 {
		return ""
	}
	b := make([]byte, n)
	copy(b, str)
	return unsafe.String(&b[0], len(b))
}

func someFunc() {
	v := RandStringRunes(1 << 10)
	fmt.Println(len(v))
	justString = Clone(v, 100)
}

func main() {
	someFunc()
	fmt.Println(len(justString))
	fmt.Println(justString)
}
