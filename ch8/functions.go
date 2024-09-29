package main

import "fmt"

func Map[T1, T2 any](s []T1, f func(T1) T2) []T2 {
	r := make([]T2, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

func Reduce[T1, T2 any](s []T1, initializer T2, f func(T2, T1) T2) T2 {
	r := initializer
	for _, v := range s {
		r = f(r, v)
	}
	return r
}

func Filter[T any](s []T, f func(T) bool) []T {
	var r []T
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

func main() {
	s := []int{1, 2, 3, 4, 5}

	mapped := Map[int, int](s, func(i int) int {
		return i * i
	})
	fmt.Println(mapped)

	reduced := Reduce[int, int](s, 1, func(i, j int) int {
		return i * j
	})
	fmt.Println(reduced)

	filtered := Filter[int](s, func(i int) bool {
		return i % 2 == 0
	})
	fmt.Println(filtered)
}
