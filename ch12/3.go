package main

import (
	"fmt"
	"math/rand"
	"time"
)

func process(v int) int {
	time.Sleep(time.Second * time.Duration(rand.Intn(10)))
	return v * 2
}

func main() {
	ch := make(chan int)
	for v := range 10 {
		go func() {
			fmt.Println(v)
			ch <- process(v)
		}()
	}
	for range 10 {
		fmt.Println("here you are", <-ch)
	}
}
