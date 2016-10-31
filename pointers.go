package main

import "fmt"

// recursion
func factorial(num int, prod int) int {
	prod = num * prod
	if num == 1 {
		return prod
	}
	return factorial(num-1, prod)
}

// accepts an int and sets it to zero
func zeroval(ival int) {

	ival = 0
}

// accepts a pointer to an integer and sets the referenced memory location to
// zero
func zeroptr(iptr *int) {

	// dereference the pointer from its memory address to the current value at
	// that address. Assigning a value to a dereferenced pointer changes the
	// value at the referenced address
	*iptr = 0
}

func add (a *int) int {
	*a = *a + 1
	return *a
}

func main() {

	// res := factorial(25, 1)
	// fmt.Println(res)
	// res = factorial(26, 1)
	// fmt.Println(res)

	fmt.Println(^uint(0))

	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("after zeroval ", i)

	fmt.Println("this is &i. It's a memory address", &i)
	zeroptr(&i)
	fmt.Println("after zeroptr: ", i)

	fmt.Println("the memory address is :", &i, ". This doesn't change as expected")
	x := 3
	fmt.Println("x = ", x)
	x1 := add(&x)

	fmt.Println("x1 is ", x1)
	fmt.Println("x is", x)

	// Passing variable pointers to function is useful when you want multiple
	// functions to make changes to the same variables faster than passing
	// variables to functions which make a copy, while a pointer is an 8 byte
	// memory address.
}
