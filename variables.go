package main

import "fmt"

func main() {

	// In Go, variables are explicitly declared and used by the compiler to e.g. check type-correctness of function calls.
	// var declares 1 or more variables
	var a string = "initial"
	fmt.Println(a)

	// decalre multiple variables at once
	var b, c int = 1, 2
	fmt.Println(b, c)

	// go has type inference
	var d = true
	fmt.Println(d)

	//variables decalred without initialization are zero valued
	var e int
	fmt.Println(e)

	// the := syntax is shorthand for decalrin and initializing a variables
	// the folllowing line is short for var f string = "short"
	f := "short"
	fmt.Println(f)
}
