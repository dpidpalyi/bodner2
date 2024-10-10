package main

import (
	"fmt"
	"sync"
)

func process(v int) int {
	return v * 2
}

func main() {
	var wg sync.WaitGroup
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ch := make(chan int)
	go func() {
		defer close(ch)
		for _, v := range arr {
			wg.Add(1)
			go func() {
				defer wg.Done()
				ch <- process(v)
			}()
		}
		wg.Wait()
	}()
	out := make([]int, 0, len(arr))
	for v := range ch {
		out = append(out, v)
	}
	fmt.Println(out)
}
