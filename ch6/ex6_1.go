package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func MakePerson(firstName string, lastName string, age int) Person {
	return Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

func MakePersonPointer(firstName string, lastName string, age int) *Person {
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

// go build -gcflags="-m" ex6_1.go
func main() {
	p1 := MakePerson("John", "Doe", 30)
	// fmt.Println receive any interface, so go compiler moves
	// to the heap any value that is passed in to a function
	// via a parameter that is of an interface type.
	fmt.Println(p1)
	p2 := MakePersonPointer("Ivan", "Ivanov", 30)
	fmt.Println(p2)
}
