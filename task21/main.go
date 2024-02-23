package main

import "fmt"

type TargetInterface interface {
	DoSomething()
}

type A struct {
	someData string
}

func (a *A) SomeMagic() {
	fmt.Println(a.someData)
}

type Adapter struct {
	*A
}

func (a *Adapter) DoSomething() {
	a.SomeMagic()
}

func NewAdapter(a *A) TargetInterface {
	return &Adapter{a}
}

func main() {
	adapter := NewAdapter(&A{someData: "data"})
	adapter.DoSomething()
}
