package main

import (
	"fmt"
	"sync"
)

func process() {
	var wg sync.WaitGroup
	in := make(chan int)
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			in <- i
		}
	}()
	go func() {
		defer wg.Done()
		for i := 10; i < 20; i++ {
			in <- i
		}
	}()
	go func() {
		wg.Wait()
		close(in)
	}()
	var wg2 sync.WaitGroup
	wg2.Add(1)
	go func() {
		defer wg2.Done()
		for v := range in {
			fmt.Println(v)
		}
	}()
	wg2.Wait()
}

func main() {
	process()
}
