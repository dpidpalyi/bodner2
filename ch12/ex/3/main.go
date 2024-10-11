package main

import (
	"fmt"
	"math"
	"sync"
)

func buildSquareRootMap() map[int]float64 {
	m := make(map[int]float64, 100000)
	for i := range 100000 {
		m[i] = math.Sqrt(float64(i))
	}
	return m
}

var squareRootMapCache = sync.OnceValue(buildSquareRootMap)

func main() {
	for i := 1000; i < 100000; i += 1000 {
		fmt.Println(i, squareRootMapCache()[i])
	}
}
