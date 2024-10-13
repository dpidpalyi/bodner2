package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	p := Person{
		Name: "John",
		Age:  31,
	}
	tmpFile, err := os.CreateTemp(os.TempDir(), "sample-")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())
	err = json.NewEncoder(tmpFile).Encode(p)
	if err != nil {
		log.Fatal(err)
	}
	err = tmpFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	tmpFile2, err := os.Open(tmpFile.Name())
	if err != nil {
		log.Fatal(err)
	}
	var p2 Person
	err = json.NewDecoder(tmpFile2).Decode(&p2)
	if err != nil {
		log.Fatal(err)
	}
	err = tmpFile2.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", p2)
}
