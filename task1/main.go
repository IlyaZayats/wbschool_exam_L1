package main

import "fmt"

type HumanInterface interface {
	do() string
}

type Human struct {
	name   string
	age    int
	gender bool
}

func (h *Human) do() string {
	return fmt.Sprintf("%v years old %s have just", h.age, h.name)
}

type Action struct {
	Human
	action string
}

func act(hi HumanInterface, action string) {
	fmt.Println(fmt.Sprintf("%s %s", hi.do(), action))
}

func main() {
	Dave := Human{name: "Dave", age: 31, gender: true}
	DaveAction := &Action{Dave, "said hello!"}
	act(DaveAction, DaveAction.action)
}
