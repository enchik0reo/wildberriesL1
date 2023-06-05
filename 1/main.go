// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской структуры
// Human (аналог наследования).
package main

import "fmt"

type Human struct {
	name string
	age  uint8
}

func (h *Human) Say() {
	fmt.Printf("My name is %s, I'm %d years old\n", h.name, h.age)
}

type Action struct {
	Human
}

func main() {
	a := Action{
		Human{
			name: "Bob",
			age:  25,
		},
	}

	a.Say()
}
