package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"strings"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	s := `
	{"name": "Fred", "age": 40}
	{"name": "Mary", "age": 21}
	{"name": "Pat", "age": 30}
	`

	var p Person
	var persons []Person
	dec := json.NewDecoder(strings.NewReader(s))
	for {
		err := dec.Decode(&p)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			log.Fatal(err)
		}
		fmt.Println(p)
		persons = append(persons, p)
	}

	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	for _, p := range persons {
		err := enc.Encode(p)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Print(b.String())
}
