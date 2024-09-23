package main

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	var p []Person
	// p := make([]Person, 0, 10_000_000)
	for i := 0; i < 10_000_000; i++ {
		p = append(p, Person{"John", "Doe", 30})
	}
}
