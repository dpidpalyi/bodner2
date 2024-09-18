package main

import (
	"fmt"
	"maps"
	"slices"
)

func main() {
	// Array declaration
	var x1 [3]int
	fmt.Println(x1)
	var x2 = [3]int{10, 20, 30}
	fmt.Println(x2)
	var x3 = [3]int{10, 20}
	fmt.Println(x3)
	var x4 = [3]int{1: 10}
	fmt.Println(x4)
	var x5 = [...]int{10, 20, 30, 40}
	fmt.Println(x5)

	// Compare arrays
	fmt.Println(x3 == x4)
	fmt.Println(x3 == [3]int{10, 20})

	// Slices declaration
	var y1 = []int{10, 20, 30}
	fmt.Println(y1)
	var y2 = []int{1, 5: 4, 6}
	fmt.Println(y2)
	// nil slice
	var y3 []int
	fmt.Println(y3)
	// slice can compare only with nil
	fmt.Println(y3 == nil)
	// we can compare with slices package
	var y4 = []int{1, 2, 3}
	var y5 = []int{1, 2, 4}
	fmt.Println(slices.Equal(y4, y5))

	// len function
	fmt.Println(len(y2))
	fmt.Println(len(y3))

	// append
	var z []int
	z = append(z, 10)
	fmt.Println(z)
	z = append(z, 20, 30, 40)
	fmt.Println(z)
	fmt.Println(len(z), cap(z))
	z = append(z, 50)
	fmt.Println(len(z), cap(z))

	s := []string{"a", "b", "c", "d"}
	s1 := s[:2]
	s2 := s[1:]
	s3 := s[1:3]
	s4 := s[:]
	fmt.Println(s)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)

	// taking slice from slice share memory
	c1 := []int{1, 2, 3, 4}
	c2 := c1[:2]
	c3 := c1[1:3]
	c3[0] = 100
	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c3)

	// copy
	xx := []int{1, 2, 3, 4}
	yy := make([]int, 3)
	num := copy(yy, xx)
	fmt.Println(yy, num)
	num = copy(yy, xx[2:])
	fmt.Println(yy, num, len(yy), cap(yy))

	// overlapping copy
	xx = []int{1, 2, 3, 4}
	copy(xx[2:], xx[:2])
	fmt.Println(xx)

	// convert from array to slice
	xArray := [4]int{1, 2, 3, 4}
	xSlice := xArray[:2]
	xSlice = append(xSlice, 300)
	fmt.Println(xArray, xSlice, cap(xArray), cap(xSlice))

	// convert from slice to array make copy of slice: we can convert only to less eq
	// length of array
	xx = []int{1, 2, 3, 4}
	zz := [2]int(xx)
	fmt.Println(zz)

	// convert from slice to pointer of array shares memory
	xx = []int{1, 2, 3, 4}
	xPtr := (*[4]int)(xx)
	xx[0] = 100
	fmt.Println(xx, *xPtr)

	// Strings
	ss1 := "Hello Даня!"
	ss2 := ss1[6:9]
	bss := []byte(ss2)
	bss[0] = 'J'
	fmt.Println(ss2)
	fmt.Println(bss)

	// Maps
	// we can't write to nil map: if u want to use it after declaaraion
	// u should assign another map to it or return from make function
	// but attempting to read values from map returns zero value
	var nilMap map[string]int
	// panic
	// nilMap["panic"] = 10
	// no panic
	fmt.Println(nilMap["no data"])
	nilMap = make(map[string]int)
	nilMap["panic"] = 10
	fmt.Println(nilMap)

	// Creating with map literal
	// zero length map
	// zeroMap := map[int]int{}
	newMap := map[string]int{
		"hi":  100,
		"bye": 200,
	}
	fmt.Println(newMap)

	// Creating with make func with size, but lenth still zero(like in slices)
	anMap := make(map[float64]int, 10)
	fmt.Println(anMap, len(anMap))

	// Like slices maps are not comparable, we can compare only with nil
	fmt.Println(nilMap == nil)

	// PS: key for a map can be any comparable type, slices can't used as keys

	// comma ok idiom with maps
	v, ok := newMap["h"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("map haven't value", v)
	}

	// deleting from map
	delete(newMap, "bye")
	// can delete from nil map and elements that doesn't exist(no panic)
	var nilMap2 map[int]int
	delete(nilMap2, 1234)
	delete(map[int]int{100: 10}, 10000)
	fmt.Println(newMap)

	// clear map or slice
	// maps deletes keys(decrease length to 0)
	// slices sets elements to zero values(keep length same)
	sss := []int{10, 20, 30, 40}
	mmm := map[int]int{10: 10, 20: 20}
	fmt.Println(sss, len(sss), mmm, len(mmm))
	clear(sss)
	clear(mmm)
	fmt.Println(sss, len(sss), mmm, len(mmm))

	// Comparing maps with 'maps' package
	m1 := map[int]int{
		10: 10,
		20: 20,
	}
	m2 := map[int]int{
		20: 20,
		10: 10,
	}
	fmt.Println(maps.Equal(m1, m2))

	// Using maps to simulate sets
	set := map[int]bool{
		1: true,
		2: true,
	}
	fmt.Println(len(set))
	if set[3] {
		fmt.Println("3 is in the set")
	} else {
		fmt.Println("3 is not in the set")
	}

	// Structs
	// We can define struct type outside and inside functions(like this main)
	type person struct {
		name string
		age  int
	}

	// Declaring variables
	// empty structs
	var liza person
	mike := person{}
	fmt.Println(liza, mike)

	// nonempty structs
	bob := person{
		"Bob",
		25,
	}

	alice := person{
		name: "Alice",
		age: 18,
	}
	fmt.Println(bob, alice)
	
	// access to fields with dot
	fmt.Println(bob.name, alice.age)

	// anonymous structs
	cat := struct{
		name string
		say string
	}{
		name: "cat",
		say: "meow",
	}
	fmt.Println(cat)

	// comparing structs (==, !=)
	fmt.Println(bob == alice)
	
	// converting structs allowing only with names, order and types
	type secondPerson struct {
		name string
		age int
	}

	second := secondPerson{"James", 30}
	fmt.Println(second)
	fmt.Println(person(second))
}
