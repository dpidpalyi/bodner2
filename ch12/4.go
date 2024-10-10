package main

import (
	"fmt"
	"time"
	"math/rand"
)

func process() int {
	v := rand.Intn(10)
	time.Sleep(time.Second * time.Duration(v))
	return v
}

func main() {
	ch := make(chan int)
	done := make(chan int)
	for range 10 {
		go func() {
			ch <- process()
		}()
	}
	go func() {
		fmt.Println("waiting to close all")
		time.Sleep(time.Second * 5)
		done <- 10
	}()
	for {
		select {
		case v := <-ch:
			fmt.Println("ch =", v)
		case <-done:
			fmt.Println("finish...")
			return
		}
	}
}
