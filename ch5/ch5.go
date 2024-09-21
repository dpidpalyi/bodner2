package main

import (
	"errors"
	"fmt"
	"os"
	"sort"
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
func divAndRemainder(num int, denom int) (int, int, error) {
	if denom == 0 {
		return 0, 0, errors.New("can't divide by zero")
	}
	return num / denom, num % denom, nil
}

// Named Return Values
// named return values are initialized to their zero values when created
func divAndRemainderNamed(num int, denom int) (result int, remainder int, err error) {
	if denom == 0 {
		err = errors.New("cannot divide by zeri")
		return result, remainder, err
	}
	result, remainder = num/denom, num%denom
	return result, remainder, err
}

// We can name only some of the return values, for other use '_'
func foo() (result int, _ int) {
	result = 5
	return result, 0
}

// We can use Blank returns, but DON'T DO THIS(NEVER)
func foo2() (result int, _ int) {
	result = 5
	return
}

func f1(a string) int {
	return len(a)
}

func f2(a string) int {
	total := 0
	for _, r := range a {
		total += int(r)
	}
	return total
}

func test(i int) {
	fmt.Println("HAHAHA", i)
}

func makeMult(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
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

	for i := 3; i >= 1; i-- {
		result, reminder, err := divAndRemainder(10, i)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(result, reminder)
	}

	a, b := foo2()
	fmt.Println(a, b)

	// Functions are values
	// type of a function is built out of keyword func and the types of the
	// parameters and return valeus => is called "signature"

	// we can declare a function variable
	var myFuncVariable func(string) int
	myFuncVariable = f1
	fmt.Println(myFuncVariable("Hello"))
	myFuncVariable = f2
	fmt.Println(myFuncVariable("Hello"))

	// The default zero value for function variable is nil
	// Attempt to run with nil value result in panic
	var nilVar func() int
	fmt.Println(nilVar)
	// panic
	// nilVar()

	// Anonymous function
	f := func(i int) {
		fmt.Println("printig", i, "from outside of an anonymous function")
	}
	for i := 0; i < 5; i++ {
		// just like any other function anonymous function is called
		// by using parentheses.
		f(i)
	}
	f = test
	f(10)

	// We don't have to assign an anonymous function to a variable before
	// use
	for i := 0; i < 5; i++ {
		func(j int) {
			fmt.Println("print", j)
		}(i)
	}

	var even = func(n int) bool {
		return n%2 == 0
	}

	fmt.Println(even(10))
	fmt.Println(even(11))

	// CLOSURES
	// functions declared inside functions ARE CLOSURES(!)
	// this meansh functions are able to access and modify varibles
	// declared in the outer function

	for i := 0; i < 5; i++ {
		func() {
			fmt.Println("print", i, "variable in the outer(main) function")
		}()
	}

	// We can shadow a varible inside a closure
	A := 20
	f1 := func() {
		fmt.Println(A)
		A := 30
		fmt.Println(A)
	}
	f1()
	fmt.Println(A)

	// Closures can reduce the number of declarations at the package level,
	// which can make it easier to find an unused name.
	// Closures really become interesting when they are passed to other
	// functions or returned from a function.

	ff := func(number int) func() int {
		return func() int {
			number *= 10
			return number
		}
	}
	omg := ff(100)
	for i := 0; i < 3; i++ {
		fmt.Println(omg())
	}

	// Passing functions as Parameters
	type Person struct {
		FirstName string
		LastName  string
		Age       int
	}

	people := []Person{
		{"Pat", "Patterson", 37},
		{"Tracy", "Bobdaughter", 23},
		{"Fred", "Fredson", 18},
	}

	// sort by last name
	sort.Slice(people, func(i, j int) bool {
		return people[i].LastName < people[j].LastName
	})
	fmt.Println(people)

	// sort by age
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println(people)

	// Returning Functions from Functions
	twoBase := makeMult(2)
	threeBase := makeMult(3)
	for i := 0; i < 5; i++ {
		fmt.Println(twoBase(i), threeBase(i))
	}

	// DEFER
	// 1. You use a function, method or closure with defer.
	// 2. Multiple defer run in last-in, first-out(LIFO) order(like stack)
	// 3. The code within defer runs AFTER the return statement. 
	// 4. The input pararmeters are evaluated immediately(!)

	hehehe := func() int {
		a := 10
		defer func(val int) {
			fmt.Println("first:", val)
		}(a)
		a = 20
		defer func(val int) {
			fmt.Println("second:", val)
		}(a)
		a = 30
		fmt.Println("exiting:", a)
		return a
	}
	fmt.Println(hehehe())

	// GO is Call by VALUE(!!!)
	type person struct {
		age int
		name string
	}

	gg := func(i int, s string, p person) {
		i *= 2
		s = "Goodbye"
		p.name = "Bob"
	}
	p := person{}
	i := 2
	s := "Hello"
	gg(i, s, p)
	fmt.Println(i, s, p)

	// Slices and maps differs

	// Any changes to a map parameter are reflected in the variable passed
	// into the function
	mm := func(m map[int]string) {
		m[2] = "hello"
		m[3] = "goodbye"
		delete(m, 1)
	}

	// You can modify any element in the slice, but you can't lengthen the slice
	ss := func(s []int) {
		for i, v := range s {
			s[i] = v * 2
		}
		s = append(s, 10)
	}
	m := map[int]string{
		1: "first",
		2: "second",
	}
	mm(m)
	fmt.Println(m)

	s1 := []int{1, 2, 3}
	ss(s1)
	fmt.Println(s1)
}
