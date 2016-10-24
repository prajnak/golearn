package main

import "fmt"

import "time"

// switch statements express conditionals across many branches

func main() {

	i := 1
	fmt.Println("write ", i, "as ")

	// each case can have more than one statement
	switch i {
	case 1:
		fmt.Println("one")
		fmt.Println("can i do this?")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// You can use commas to separate multiple expressions in the same case
	// statement. We use the optional default case in this example as well.
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("weekend")
	default:
		fmt.Println("weekdaysss")
	}

	// switch without an expression is an alternate way to express if/else
	// logic. Here we also show how the case expressions can be non-constants.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("it's before noon")
	default:
		fmt.Println("after noon")
	}
}
