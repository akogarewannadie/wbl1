package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h *Human) Speak() {
	fmt.Printf("%s says hello\n", h.Name)
}

type Action struct {
	Human
}

func main() {

	action := Action{
		Human: Human{Name: "John", Age: 30},
	}

	action.Speak()
}
