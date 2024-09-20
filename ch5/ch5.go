package main

import (
	"errors"
	"fmt"
	"os"
)

// functions
func div(num int, denom int) int {
	if denom == 0 {
		return 0
	}
	return num / denom
}

// Simulating Named and Optional Parameters
type MyFuncOpts struct {
	FirstName string
	LastName  string
	Age       int
}

// Variadic Input Parameters And slices
// variadic must be last or only parameter
func Foo(mul int, vals ...int) {
	for _, v := range vals {
		fmt.Println(mul * v)
	}
}

func MyFunc(opts MyFuncOpts) error {
	if opts.FirstName == "" {
		return errors.New("missed FirstName")
	}
	fmt.Println("FirstName:", opts.FirstName)
	if opts.LastName != "" {
		fmt.Println("LastName:", opts.LastName)
	}
	if opts.Age != 0 {
		fmt.Println("Age:", opts.Age)
	}

	return nil
}

// Multiple return values
func divAndReminder(num int, denom int) (int, int, error) {
	if denom == 0 {
		return 0, 0, errors.New("can't divide by zero")
	}
	return num / denom, num % denom, nil
}

func main() {
	fmt.Println(div(10, 2))
	err := MyFunc(MyFuncOpts{FirstName: "John"})
	if err != nil {
		fmt.Println(err)
	}
	err = MyFunc(MyFuncOpts{})
	if err != nil {
		fmt.Println(err)
	}
	Foo(10, 1, 2, 3, 4)
	Foo(10, []int{1, 2, 3, 4}...)

	for i := 3; i >= 0; i-- {
		result, reminder, err := divAndReminder(10, i)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(result, reminder)
	}
}
