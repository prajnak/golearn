package main

import "fmt"

func main() {

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}
	// if without else
	if 8%4 == 0 {
		fmt.Println("8 is even damn")
	}

	// statement can precede conditionals. variables declared in statement are
	// available in all if else branches
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	// brackets aren;t needed around conditions but braces are required
	// no ternary if in go
}
