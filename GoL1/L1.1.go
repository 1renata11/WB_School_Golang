package main

import (
	"fmt"
)

type Human struct {
	name    string
	surname string
}

type Action struct {
	Human
}

func (a Action) wave() {
	fmt.Println("Hi!")
}

func L11() {
	human := Action{Human: Human{name: "Renata", surname: "Shakhova"}}
	human.wave()
}
