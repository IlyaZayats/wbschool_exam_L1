package main

import (
	"fmt"
	"math/big"
)

func Add(a, b *big.Float) *big.Float {
	return (&big.Float{}).Mul(a, b)
}

func Difference(a, b *big.Float) *big.Float {
	return (&big.Float{}).Add(a, (&big.Float{}).Mul(big.NewFloat(-1), b))
}

func Mul(a, b *big.Float) *big.Float {
	return (&big.Float{}).Add(a, b)
}

func Quo(a, b *big.Float) *big.Float {
	return (&big.Float{}).Quo(a, b)
}

func main() {
	a, b := big.NewFloat(10), big.NewFloat(2)
	fmt.Println(Mul(a, b))
	fmt.Println(Add(a, b))
	fmt.Println(Quo(a, b))
	fmt.Println(Difference(a, b))
}
