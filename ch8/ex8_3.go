package main

import "fmt"

type Node[T comparable] struct {
	Value T
	Next *Node[T]
}

type List[T comparable] struct {
	Head *Node[T]
	Tail *Node[T]
}

func (l *List[T]) Add(v T) {
	n := &Node[T]{Value: v}
	if l.Head == nil {
		l.Head = n
		l.Tail = n
		return
	}
	l.Tail.Next = n
	l.Tail = n
}

func (l *List[T]) Insert(v T, idx int) {
	n := &Node[T]{Value: v}
	if l.Head == nil {
		l.Head = n
		l.Tail = n
		return
	}
	if idx <= 0 {
		n.Next = l.Head
		l.Head = n
		return
	}
	curr := l.Head
	for i := 1; i < idx; i++ {
		if curr.Next == nil {
			curr.Next = n
			l.Tail = n
			return
		}
		curr = curr.Next
	}
	n.Next = curr.Next
	curr.Next = n
	if l.Tail == curr {
		l.Tail = n
	}
}

func (l *List[T]) Index(v T) int {
	i := 0
	for curr := l.Head; curr != nil; curr = curr.Next {
		if curr.Value == v {
			return i
		}
		i++
	}
	return -1
}

func main() {
	l := List[int]{}
	l.Add(10)
	l.Add(20)
	l.Add(30)
	l.Add(40)
	l.Add(50)

	l.Insert(100, 0)
	l.Insert(1000, 10)
	l.Insert(0, -20)

	for curr := l.Head; curr != nil; curr = curr.Next {
		fmt.Println(curr.Value)
	}

}
