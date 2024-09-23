package main

import (
	"fmt"
	"time"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// Method
// Go support methods on user-defined(!) types.
// NB!!
// Methods can be define only at the package block level, while
// functions can be defined inside any block
// Functions and method names cannot be overloaded.
// Keep methods declared in the same package as their associated type

// p is a reciever
func (p Person) String() string {
	return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
}

// Pointer Receivers And Value Receivers
// If method modifies the receiver, you MUST use a pointer receiver
// If method needs to handle nil instances, then it MUST use a pointer receiver
// If method doesn't modify receiver, you can use a value receiver.
// Don't mix both!

type Counter struct {
	total       int
	lastUpdated time.Time
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}

// User-defined types
type Score int
type Converter func(string) Score
type TeamScores map[string]Score

// METHODS ARE FUNCTIONS TOO
type Adder struct {
	start int
}

func (a Adder) AddTo(val int) int {
	return a.start + val
}

func main() {
	//var t = TeamScores{}
	t := make(TeamScores)
	t["Hi"] = 10
	fmt.Println(t)
	var c Counter
	fmt.Println(c.String())
	c.Increment()
	fmt.Println(c.String())
	ptrC := &Counter{}
	fmt.Println(ptrC.String())
	ptrC.Increment()
	fmt.Println(ptrC.String())

	// If you call a value receiver method with pointer instance
	// whose value is nil, code will compile but will panic at runtime
	//var nilPtr *Counter
	// panic
	//nilPtr.Increment()
	//fmt.Println(nilPtr.String())

	// The rulse for passing values to functions still apply to methods
	// They call-by-value

	// Methods like a functions
	myAdder := Adder{start: 10}
	fmt.Println(myAdder.AddTo(5))
	type adder func(int) int
	var f1 adder
	// This is called a METHOD VALUE
	f1 = myAdder.AddTo
	fmt.Println(f1(10))
	// You can create a function from the type itself
	// This is called a METHOD EXPRESSION
	f2 := Adder.AddTo
	fmt.Println(f2(myAdder, 20))
	// One way to use them at dependency injection

	// IOTA
	// the best practice it to first define a type based on int
	type MailCategory int
	// next use a const block to define a set of values for your type
	const (
		Uncategorized MailCategory = iota // 0
		Personal			  // 1
		Spam				  // 2
		Social				  // 3
		Advertisements			  // 4
	)
	// NB!
	// When new const block is created, iota is set back to 0
}
