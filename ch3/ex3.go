package main

import "fmt"

func main() {
	// 1.
	greetings := []string{"Hello", "Hola", "←→", "abc", "Привет"}
	s1 := greetings[:2]
	s2 := greetings[1:4]
	s3 := greetings[3:]
	fmt.Println(greetings)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)

	// 2.
	message := "Hi 😂 and 😂"
	r := []rune(message)
	fmt.Println(string(r[3]))
	
	// 3.
	type Employee struct {
		firstName string
		lastName string
		id int
	}

	e1 := Employee{"John", "Doe", 1}
	e2 := Employee{
		firstName: "Alex",
		lastName: "Edwards",
		id: 2,
	}

	var e3 Employee
	e3.firstName = "Jon"
	e3.lastName = "Bodner"
	e3.id = 3
	fmt.Println(e1)
	fmt.Println(e2)
	fmt.Println(e3)
}
