package main

import (
	"fmt"
	"slices"
)

func main() {
	// Array declaration
	var x1 [3]int
	fmt.Println(x1)
	var x2 = [3]int{10, 20, 30}
	fmt.Println(x2)
	var x3 = [3]int{10, 20}
	fmt.Println(x3)
	var x4 = [3]int{1:10}
	fmt.Println(x4)
	var x5 = [...]int{10, 20, 30, 40}
	fmt.Println(x5)

	// Compare arrays
	fmt.Println(x3 == x4)
	fmt.Println(x3 == [3]int{10, 20})

	// Slices declaration
	var y1 = []int{10, 20, 30}
	fmt.Println(y1)
	var y2 = []int{1, 5: 4, 6}
	fmt.Println(y2)
	// nil slice
	var y3 []int
	fmt.Println(y3)
	// slice can compare only with nil
	fmt.Println(y3 == nil)
	// we can compare with slices package
	var y4 = []int{1, 2, 3}
	var y5 = []int{1, 2, 4}
	fmt.Println(slices.Equal(y4, y5))

	// len function
	fmt.Println(len(y2))
	fmt.Println(len(y3))

	// append
	var z []int
	z = append(z, 10)
	fmt.Println(z)
	z = append(z, 20, 30, 40)
	fmt.Println(z)
	fmt.Println(len(z), cap(z))
	z = append(z, 50)
	fmt.Println(len(z), cap(z))

	s := []string{"a", "b", "c", "d"}
	s1 := s[:2]
	s2 := s[1:]
	s3 := s[1:3]
	s4 := s[:]
	fmt.Println(s)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)

	// taking slice from slice share memory
	c1 := []int{1, 2, 3, 4}
	c2 := c1[:2]
	c3 := c1[1:3]
	c3[0] = 100
	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c3)

	// copy 
	xx := []int{1, 2, 3, 4}
	yy := make([]int, 3)
	num := copy(yy, xx)
	fmt.Println(yy, num)
	num = copy(yy, xx[2:])
	fmt.Println(yy, num, len(yy), cap(yy))
}
