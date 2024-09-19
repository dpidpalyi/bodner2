package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// 1.
	sl := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		sl = append(sl, rand.Intn(100))
	}
	fmt.Println(sl)

	// 2.
	for _, v := range sl {
		switch {
		case v % 2 == 0 && v % 3 == 0:
			fmt.Println("Six!")
		case v % 2 == 0:
			fmt.Println("Two!")
		case v % 3 == 0:
			fmt.Println("Three!")
		default:
			fmt.Println("Never mind")
		}
	}

	// 3.
	var total int
	for i := 0; i < 10; i++ {
		// it prints 0 .. 9
		// the bug is shadowing upper total variable
		//total = total + i
		total := total + i
		fmt.Println(total)
	}
}
