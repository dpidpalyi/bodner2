package main

import "fmt"

const (
	Field1 = 10
	Field2
	Field3 = iota
)

func main() {
	fmt.Println(Field1, Field2, Field3)
}
