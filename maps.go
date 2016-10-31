package main

import "fmt"

func main() {

	// maps can be made using the make command
	m := make(map[string]int)
	fmt.Println(m)
	m["k1"] = 7
	m["k2"] = 9

	// They can also be made without the make command like
	maap := map[string] int{"C": 1, "B": 3}
	fmt.Println(maap)
	// fmt.Println(m, m["k1"], len(m))

	// //delete using delete builtin

	// delete(m, "k1")
	// fmt.Println(m)

	// // access deleted key
	// mm, prs := m["k1"]
	// fmt.Println(mm, prs)

	// make is used for memory allocation of the builtin models like map, slice
	// and channel new is for types memory allocation new(T) allocates zero
	// value to type T's memory, returns its memory address, which is the value
	// of type *T. Returns a pointer which points to the type T's zero value

	// make(T, args) returns a type T with an initial value.
}
