package main

import "fmt"

func processChannel(ch <-chan int) []int {
	const conc = 10
	results := make(chan int, conc)
	for i := 0; i < conc; i++ {
		go func() {
			v := <-ch
			results <- process(v)
		}()
	}
	out := make([]int, 0, conc)
	for i := 0; i < conc; i++ {
		out = append(out, <-results)
	}
	return out
}

func process(v int) int {
	return v * 2
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	ch := make(chan int)
	go func() {
		for _, v := range arr {
			ch <- v
		}
	}()
	result := processChannel(ch)
	fmt.Println(result)
}
