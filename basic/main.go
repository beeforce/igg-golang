package main

import "fmt"

type person struct {
	firstName string
	lastName string
}

func (p person) Print() {
	fmt.Println(p)
}

func main() {
	bee := person{
		"Buri",
		"Paoton",
	}

	bee.Print()
}