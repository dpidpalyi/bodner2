package main

import "fmt"

const value = 100

func main() {
	// 1.
	i := 20
	f := float64(i)
	fmt.Printf("i of type %T: %d\n", i, i)
	fmt.Printf("f of type %T: %.3f\n", f, f)

	// 2.
	i = value
	f = value
	fmt.Printf("i of type %T: %d\n", i, i)
	fmt.Printf("f of type %T: %.3f\n", f, f)

	// 3.
	var (
		b      byte   = 1 << 8 - 1
		smallI int32  = 1 << 31 - 1
		bigI   uint64 = 1 << 64 - 1
	)
	b += 1
	smallI += 1
	bigI += 1
	fmt.Printf("b of type %T: %d\n", b, b)
	fmt.Printf("smallI of type %T: %d\n", smallI, smallI)
	fmt.Printf("bigI of type %T: %d\n", bigI, bigI)
}
