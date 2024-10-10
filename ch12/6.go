package main

import "fmt"

func countTo(n int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := range n {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func main() {
	for v := range countTo(10) {
		if v > 5 {
			break
		}
		fmt.Println(v)
	}
}
