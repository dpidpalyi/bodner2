package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"strings"
)

func countLetters(r io.Reader) (map[string]int, error) {
	buf := make([]byte, 2048)
	out := map[string]int{}
	for {
		n, err := r.Read(buf)
		for _, b := range buf[:n] {
			if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
				out[string(b)]++
			}
		}
		if err == io.EOF {
			return out, nil
		}
		if err != nil {
			return nil, err
		}
	}
}

func buildGZipReader(filename string) (*gzip.Reader, func(), error) {
	r, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	gr, err := gzip.NewReader(r)
	if err != nil {
		return nil, nil, err
	}
	return gr, func() {
		gr.Close()
		r.Close()
	}, nil
}

func main() {
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	out, err := countLetters(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(out)
	sr := strings.NewReader("Hello world")
	out, err = countLetters(sr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(out)

	r, closer, err := buildGZipReader("file.txt.gz")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer closer()
	counts, err := countLetters(r)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(counts)
}
