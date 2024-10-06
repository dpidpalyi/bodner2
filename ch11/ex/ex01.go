package main

import (
	_ "embed"
	"fmt"
	"os"
)

//go:embed english_rights.txt
var eng string

//go:embed hungarian_rights.txt
var hung string

//go:embed italian_rights.txt
var ita string

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}
	switch os.Args[1] {
	case "english":
		fmt.Println(eng)
	case "hungarian":
		fmt.Println(hung)
	case "italian":
		fmt.Println(ita)
	default:
		fmt.Println("unknown language:", os.Args[1])
	}
}
