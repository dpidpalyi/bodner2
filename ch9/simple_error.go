package main

import (
	//"errors"
	"fmt"
)

func doubleEven(i int) (int, error) {
	if i % 2 != 0 {
		//return 0, errors.New("only even numbers are processed")
		return 0, fmt.Errorf("%d isn't an even number", i)
	}
	return i * 2, nil
}

func main() {
	result, err := doubleEven(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
