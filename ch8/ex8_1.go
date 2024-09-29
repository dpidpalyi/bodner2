package main

import "fmt"

type Numberer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
	~float32 | ~float64
}

func Doubler[T Numberer](n T) T {
	return n * 2
}

func main() {
	fmt.Println(Doubler(-10))
	fmt.Println(Doubler(byte(20)))
	fmt.Println(Doubler(3.14))
}
