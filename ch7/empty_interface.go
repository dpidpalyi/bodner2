package main

import "fmt"

type I int

func (i I) String() {
	fmt.Println(i)
}

func main() {
	// Empty interface type states that the variable can store any
	// value whose type implements zero or more methods.
	var i interface{}
	i = 20
	i = "hello"
	i = struct {
		FirstName string
		LastName string
	}{"Fred", "Fredson"}
	fmt.Println(i)
	i = I(100)
	fmt.Println(i)
}
