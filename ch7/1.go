package main

import "fmt"

type Low int

func (l Low) Print() {
	fmt.Println(l)
}

type High Low

func main() {
	var i int = 100
	var l Low = 200
	var h High = 300
	l = Low(i)
	h = High(i)
	fmt.Println(l, h)
	l.Print()
	h.Print()
}
