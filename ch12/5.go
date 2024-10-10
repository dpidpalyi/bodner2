package main

import (
	"context"
	"fmt"
)

func countTo(ctx context.Context, n int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := range n {
			select {
			case <-ctx.Done():
				return
			case ch <- i:
			}
		}
	}()
	return ch
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for v := range countTo(ctx, 10) {
		if v > 5 {
			break
		}
		fmt.Println(v)
	}
}
