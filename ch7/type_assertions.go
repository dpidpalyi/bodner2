package main

import "fmt"

func main() {
	//typeAssert()
	//typeAssertPanicWrongType()
	//typeAssertPanicTypeNotIdentical()
	//err := typeAssertCommaOK()
	//if err != nil {
	//	fmt.Println(err)
	//}
	var i any
	a := true
	i = a
	doThings(i)
	i = nil
	doThings(i)
}

func typeAssertCommaOK() error {
	var i any
	var mine MyInt = 20
	i = mine
	i2, ok := i.(int)
	if !ok {
		// we are constructing a new error with fmt.Errorf.
		// fmt.Errorf is covered in chapter 9.
		return fmt.Errorf("unexpected type for %v", i)
	}
	fmt.Println(i2 + 1)
	return nil
}

func typeAssertPanicTypeNotIdentical() {
	// we are using recover to allow us to run through the
	// failing type assertions. recover is explained in chapter 9.
	defer func() {
		if m := recover(); m != nil {
			fmt.Println(m) // prints out because a panic happens
		}
	}()
	var i any
	var mine MyInt = 20
	i = mine
	i2 := i.(int)
	fmt.Println(i2 + 1)
}

func typeAssertPanicWrongType() {
	// we are using recover to allow us to run through the
	// failing type assertions. recover is explained in chapter 9.
	defer func() {
		if m := recover(); m != nil {
			fmt.Println(m) // prints out because a panic happens
		}
	}()
	var i any
	var mine MyInt = 20
	i = mine
	i2 := i.(string)
	fmt.Println(i2)
}

type MyInt int

func doThings(i any) {
	switch j := i.(type) {
	case nil:
		fmt.Println("nil interface")
	case string:
		fmt.Println("string")
	case int:
		fmt.Println("int")
	case MyInt:
		fmt.Println("MyInt")
	case bool:
		fmt.Println(j)
	}
}

func typeAssert() {
	var i any
	var mine MyInt = 20
	i = mine
	i2 := i.(MyInt)
	fmt.Println(i2 + 1)
}
