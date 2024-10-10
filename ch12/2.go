package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		inG := 1
		ch1 <- inG
		fromM := <-ch2
		fmt.Println("in goroutine:", inG, fromM)
	}()
	inM := 2
	var fromG int
	select {
	case ch2 <- inM:
	case fromG = <-ch1:
	}
	fmt.Println("in main:", inM, fromG)
}
