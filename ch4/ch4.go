package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Shadowing variables
	x := 5
	if 1 < 2 {
		fmt.Println(x)
		// := reuses only variables declared at current block!!
		x, y := 10, 20
		fmt.Println(x, y)
	}
	fmt.Println(x)

	// keywords like int, string, true, false, make and close are
	// in universe block, that contains all other blocks!

	// if statement with variable 'i' at if block
	if i := rand.Intn(10); i < 5 {
		fmt.Println("low")
	} else if i > 5 {
		fmt.Println("high")
	} else {
		fmt.Println("mid")
	}

	// for
	// c-style for
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	// condition only
	j := 5
	for j > 0 {
		fmt.Println(j)
		j--
	}
	// infinite for
	for {
		fmt.Println(j)
		if j == 5 {
			break
		}
		j++
	}

	// idiomatic using continue
	for i := 1; i < 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
			continue
		}
		if i%3 == 0 {
			fmt.Println("Fizz")
			continue
		}
		if i%5 == 0 {
			fmt.Println("Buzz")
			continue
		}
		fmt.Println(i)
	}

	// for-range only for built-in compound types
	for i, v := range []int{10, 20, 30} {
		fmt.Println(i, v)
	}
	// with underscore(_) to ignore range variables
	for _, v := range []string{"Hi", "Hola", "Ку"} {
		fmt.Println(v)
	}
	// only index
	for i := range []int{1, 2, 3} {
		fmt.Println(i)
	}

	// new range only with ints
	for range 5 {
		fmt.Println("oh")
	}

	// iterating over maps - always different order
	// when just print map - same order as input
	m := map[int]int{1: 10, 2: 20, 3: 30}
	for i := 0; i < 3; i++ {
		for k, v := range m {
			fmt.Println(k, v)
		}
		fmt.Println(m)
	}

	// iterating over strings gets runes
	samples := []string{"hello", "hola", "привет"}
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
		}
		fmt.Println()
	}

	// test with invalid UTF-7 values
	//b := []byte{194, 163}
	b := []byte{0b11110011, 0b01100110}
	fmt.Printf("%b %b %b\n", b[0], b[1], int32(127774))
	s := string(b)
	for i, r := range s {
		fmt.Println(i, r, string(r))
	}

	// for-range value is a copy with new temp var on each iteration
	arr := []int{10, 20, 30}
	for _, v := range arr {
		v *= 10
	}
	fmt.Println(arr)

	// labeled for (outer)
	idx := []int{10, 20, 30, 40, 50}
	for i := 0; i < 5; i++ {
	outer:
		for j := range idx {
			if i%2 == 0 {
				continue outer
			}
			fmt.Print(idx[j], " ")
		}
		fmt.Println()
	}

	// switch expression
	words := []string{"a", "cow", "smile", "gopher", "octopusese"}
	for _, word := range words {
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Println(word, "is a short word!")
		case 5:
			fmt.Println(word, "is exactly the right length 5")
		case 6, 7, 8, 9:
		default:
			fmt.Println(word, "is a long word!")
		}
	}

	// simplest switch without variable "FizzBuzz"
	// NB!
	// favor blank switch statements over if/else chains when you
	// have multiple RELATED(!) cases.
	for i := 1; i < 16; i++ {
		switch {
		case i % 3 == 0 && i % 5 == 0:
			fmt.Println("FizzBuzz")
		case i % 3 == 0:
			fmt.Println("Fizz")
		case i % 5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}

	// usefull for label with switch
	loop:
	for i := 0; i < 10; i++ {
		switch i {
		case 0, 2, 4, 6:
			fmt.Println(i, "even")
		case 1, 3, 5, 7:
			fmt.Println(i, "odd")
		case 8:
			fmt.Println("exit loop")
			break loop
		default:
			fmt.Println("doesn't matter")
		}
	}
			
}
