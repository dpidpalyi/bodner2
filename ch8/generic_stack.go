package main

import "fmt"

//type Stack[T any] struct {
type Stack[T comparable] struct {
	vals []T
}

func (s *Stack[T]) Push(val T) {
	s.vals = append(s.vals, val)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.vals) == 0 {
		var zero T
		return zero, false
	}
	top := s.vals[len(s.vals) - 1]
	s.vals = s.vals[:len(s.vals) - 1]
	return top, true
}

func (s Stack[T]) Contains(val T) bool {
	for _, v := range s.vals {
		if v == val {
			return true
		}
	}
	return false
}

func main() {
	s := Stack[int]{}
	s.Push(10)
	s.Push(20)
	s.Push(30)
	fmt.Println(s.Contains(20))
	fmt.Println(s.Contains(40))
	for range 4 {
		top, ok := s.Pop()
		if !ok {
			fmt.Println("empty stack")
			break
		}
		fmt.Println(top)
	}
}
