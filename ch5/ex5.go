package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func add(i int, j int) (int, error) {
	return i + j, nil
}

func sub(i int, j int) (int, error) {
	return i - j, nil
}

func mul(i int, j int) (int, error) {
	return i * j, nil
}

func div(i int, j int) (int, error) {
	if j == 0 {
		return 0, errors.New("can't divide by zero")
	}
	return i / j, nil
}

var opMap = map[string]func(int, int) (int, error){
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}

func fileLen(path string) (int, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	var result int
	buf := make([]byte, 2048)
	for {
		n, err := f.Read(buf)
		result += n
		if err != nil {
			if errors.Is(err, io.EOF) {
				return result, nil
			}
			return 0, err
		}
	}
}

func prefixer(s string) func(string) string {
	return func(prefix string) string {
		return s + " " + prefix
	}
}

func main() {
	// 1.
	exprs := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"2", "/", "0"},
		{"two", "+", "three"},
		{"2"},
	}

	for _, exp := range exprs {
		if len(exp) < 3 {
			fmt.Println("not enough args")
			continue
		}
		p1, err := strconv.Atoi(exp[0])
		if err != nil {
			fmt.Println(err)
			continue
		}
		op := exp[1]
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Println("unrecognized operator:", op)
			continue
		}
		p2, err := strconv.Atoi(exp[2])
		if err != nil {
			fmt.Println(err)
			continue
		}
		result, err := opFunc(p1, p2)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(result)
	}

	// 2.
	if len(os.Args) < 2 {
		log.Fatal("no file specified")
	}
	result, err := fileLen(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)

	// 3.
	helloPrefix := prefixer("Hello")
	fmt.Println(helloPrefix("Bob"))
	fmt.Println(helloPrefix("Maria"))
}
