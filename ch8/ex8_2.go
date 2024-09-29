package main

import "fmt"

type Printable interface {
	~int | ~float64
	fmt.Stringer
}

type MyInt int

func (m MyInt) String() string {
	return fmt.Sprintf("%d", m)
}

type MyFloat float64

func (m MyFloat) String() string {
	return fmt.Sprintf("%f", m)
}

func Print[T Printable](n T) {
	fmt.Println(n)
}

func main() {
	Print(MyInt(10))
	Print(MyFloat(3.14))
}
