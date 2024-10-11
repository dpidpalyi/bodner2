package main

import (
	"fmt"
	"sync"
)

func process() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	go func() {
		defer wg.Done()
		for i := 10; i < 20; i++ {
			ch2 <- i
		}
		close(ch2)
	}()
	done := make(chan int)
	go func() {
		wg.Wait()
		done <- 1
	}()
	for {
		select {
		case v, ok := <-ch1:
			if !ok {
				ch1 = nil
			} else {
				fmt.Println("first:", v)
			}
		case v, ok := <-ch2:
			if !ok {
				ch2 = nil
			} else {
				fmt.Println("second:", v)
			}
		case <-done:
			return
		}
	}
}

func main() {
	process()
}
