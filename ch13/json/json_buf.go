package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	p := Person{
		Name: "John Doe",
		Age:  31,
	}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(buf.String())
}
