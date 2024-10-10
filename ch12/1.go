package main

import "fmt"

func foo1(ch chan<- int) {
	for i := range 5 {
		ch <- i
	}
	//close(ch)
}

func foo2(ch <-chan int) {
	var i int
	for v := range ch {
		fmt.Println(v)
		if i == 3 {
			return
		}
		i++
	}
}

func main() {
	ch := make(chan int)
	go foo1(ch)
	foo2(ch)
	close(ch)
}
