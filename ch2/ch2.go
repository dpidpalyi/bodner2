package main

import "fmt"

// We can assign literals to own types too
type Myint int

// No need to read variables at package level
var HAHAHA = 300

func main() {
	// Different methods of int literals: decimal, octal, bit and hex
	a := 10
	b := 0o77
	c := 0b110
	d := 0xff
	fmt.Println(a, b, c, d)

	// Bad practice to use int32 for runes(chars)
	var ac int32 = 65

	// Different methods of rune literals: decimal, two 8-bit digits, 16-bit and 32-bit digits
	bc := '\141'
	cc := '\x41'
	dc := '\u0061'
	ec := '\U00000061'
	fmt.Printf("%c %c %c %c %c\n", ac, bc, cc, dc, ec)

	// Simple string declaration with double quotes
	s := "Hello Friend\n"
	fmt.Print(s)

	// String declaration with backquote
	s1 := `Russian МАФИЯ`
	fmt.Println(s1)

	// Using literals to declare variables
	var aA Myint = 10
	fmt.Println(aA)

	// Untyped and typed const values: they are just named literals
	const l = 10
	fmt.Println(l)
	var zz byte = l
	fmt.Println(zz)

	// Need explicit conversion
	const typedL int = 100
	var LL = float64(typedL)
	fmt.Println(LL)
}
