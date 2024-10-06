// Package
package main

import (
	"errors"
	"fmt"
)

func returnErr(b bool) (err error) {
	if b {
		return errors.New("error")
	}
	return err
}

func main() {
	err := returnErr(false)
	if err != nil {
		fmt.Println(err)
	}
	err = returnErr(true)
	fmt.Println("end of program")
}
