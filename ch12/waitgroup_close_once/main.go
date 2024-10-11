package main

import (
	"fmt"
	//"math/rand"
	"sync"
	"time"
)

func processAndGather[T, R any](in <-chan T, process func(T) R, num int) []R {
	out := make(chan R, num)
	var wg sync.WaitGroup
	wg.Add(num)
	for range num {
		go func() {
			defer wg.Done()
			for v := range in {
				out <- process(v)
			}
		}()
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	var result []R
	for v := range out {
		result = append(result, v)
	}
	return result
}

func process(n int) float64 {
	time.Sleep(time.Second)
	return float64(n) * 3.3333
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	in := make(chan int)
	go func() {
		for _, v := range arr {
			in <- v
		}
		close(in)
	}()
	result := processAndGather(in, process, 1000)
	fmt.Println(result)
}
