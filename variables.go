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

	// Constants can be integers, booleans or strings
	const Pi = 3.14159
	const prefix = "blah__"

	// Integer Types include signed and unsigned types. int and uint have
	// the same length but specific length depends on the OS being used
	// var aa int8
	// var bb int32
	// c:= a+b //can't do this because of mismatched integer types

	//Floats have two types - float32 and float64 complex numbers are
	// supported as well complex128 - 64 bit real and 64 bit imaginary
	var cd complex64 = 5+5i
	fmt.Printf("Value is : %v", cd)

	// Strings
	var s string = "hello"
	// s[0] = 'c' //can't manipulate string types by index

	ccc := []byte(s)
	ccc[0] = 'c'
	s2 := string(ccc)
	fmt.Println(s2, ccc)

}
