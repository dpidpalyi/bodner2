package main

import (
	"fmt"
	"unsafe"
)

func makePointer[T any](t T) *T {
	return &t
}

type Person struct {
	Name string
}

func foo(p **Person) {
	*p = &Person{"New"}
}

func failed(p *int) {
	z := 10
	fmt.Printf("%p\n", p)
	p = &z
	fmt.Printf("%p\n", p)
}

// Don't do this
func MakeFoo(p *Person) error {
	p.Name = "John"
	return nil
}

// Do this 
func MakeFoo2() (Person, error) {
	p := Person{
		Name: "NAME",
	}
	return p, nil
}

func main() {
	// Modern computers use 8 bytes to store a pointer.
	var i *int
	fmt.Println(unsafe.Sizeof(i))

	// & is the address operator
	x := "hello"
	ptrX := &x

	// * is the indirection(dereference) operator
	fmt.Println(*ptrX)
	// trying to dereference nil pointer get panic
	//fmt.Println(*i)

	// Pointer type can be base on any type
	y := 10
	var ptrY *int
	ptrY = &y
	fmt.Println(ptrY)

	z := func() {}
	ptrZ := &z
	fmt.Println(ptrZ)

	// new creates a pointer variable and returns pointer
	// to a zero-value instance of the provided type
	var n = new(int)
	fmt.Println(n == nil) // false
	fmt.Println(*n)       // 0

	// For structs use & before a struct literal
	// You can't use an & before a primitive literal(numbers, booleans,
	// and strings) or a constant, because they don't have memory addr.
	// For primitive type declare a variable and point to it
	type Foo struct {
		Age int
		Name *string
	}

	// If we have a struct with a field of a pointer to a primitive
	// type, we can't assing a literal directly to the field.
	// first option to use a variable to hold the constant value.
	XX := "Hello"
	st := &Foo{10, &XX}
	var yy string
	ptrYY := &yy
	// But we need to be careful not to change this variable (XX)
	fmt.Println(*st.Name, *ptrYY)
	// second option is to write a generic helper function makePointer
	p := Foo{
		Age: 40,
		Name: makePointer("Samuel"),
	}
	fmt.Println(*p.Name)

	// Pointers indicate mutable parameters
	pp := &Person{"Grant"}
	fmt.Printf("%p\n", pp)
	fmt.Println(*pp)
	foo(&pp)
	fmt.Printf("%p\n", pp)
	fmt.Println(*pp)

	var zzz *int
	failed(zzz)
	fmt.Println(zzz)

	// Pointers are a Last Resort
	// Rather than populating a struct by passing a pointer to it into
	// a function, have the function instantiate and return the struct

	// The only time we should use pointer parameters to modify variable
	// is when the function expects an interface. (JSON example)
	//f := struct {
	//	Name string `json:"name"`
	//	Age int `json:"age"`
	//}{}
	//err := json.Unmarshal([]byte(`{"name": "Bob", "age": 30}`), &f)

	// Zero Value Versus No Value
	// Go pointers can be used to indicate the difference between a 
	// variable or field that's been assigned the zero value and a
	// variable or field that hasn't been assigned a value at all.
	// NB!! When not working with JSON (or other external protocols),
	// resist the temptation to use a pointer field to indicate no value.

	// The difference Between Maps and Slices
	// On an API-design level, maps are a bad choice
	// Go is a strongly typed language, rather than passing a map around,
	// use a Struct.

	// Slice that's passed to a function can have its contents modified,
	// but the slice can't be resized.
	// This makes slices ideal for reusable buffers.

}
