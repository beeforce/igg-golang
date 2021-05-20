package main

import "fmt"


type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contact contactInfo
}


func (p person) Print() {
	fmt.Println(p)
}

func main() {
	bee := person{
		"Buri",
		"Paoton",
		contactInfo{
			email: "buri@igeargeek.com",
			zipCode: 50210,
		},
	}

	bee.Print()
}