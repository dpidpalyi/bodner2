package main

import "fmt"

func UpdateSlice(s []string, str string) {
	if len(s) > 0 {
		s[len(s)-1] = str
	}
	fmt.Println(s)
}

func GrowSlice(s []string, str string) {
	s = append(s, str)
	fmt.Println(s)
}

func main() {
	s1 := []string{"Hello", "world"}
	fmt.Println(s1, "=", `[Hello world]`)
	UpdateSlice(s1, "WORLD")
	// Because when we pass slice to function, if length of slice
	// isn't change, we can change the contents of a slice
	fmt.Println(s1, "=", `[Hello WORLD]`)

	s2 := []string{"Buy", "world"}
	fmt.Println(s2, `[Buy world]`)
	GrowSlice(s2, "AHAHAHA")
	// We change the length inside function, but length of passed
	// slice doesn't changes
	fmt.Println(s2, `[Buy world]`)
}
