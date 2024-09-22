package main

import "fmt"

func foo(s []int) {
	s[0] = 100
	//s = append(s, 200, 300, 400)
	s = append(s, 200, 300)
	s[0] = 100000
}

func main() {
	s := make([]int, 0, 4)
	s = append(s, 1, 2)
	foo(s)
	fmt.Println(s)
	s = s[:4]
	fmt.Println(s)
}
