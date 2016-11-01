package main

import (
	"fmt"
)

// go supports concurrency at the language level goroutines and concurrency are
// core go concepts Goroutines are similar to threads but work differently A
// dozen goroutines can have 5-6 underlying threads and have full support for
// sharing memory between goroutines. Each goroutine takes 4-5kb of stack memory
// Goroutines are more lightweight, efficient and convenient than system threads

//goroutines run on the thread manager at runtime in Go. Use the "go" keyword to
//create a new goroutine, which is a function at the underlying level
// aside - main() is a goroutine

func say(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
}

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// call the function f synchronously
	f("direct")

	// invoke f in a goroutine which executes concurrently with the calling one
	// (because main() is also a goroutine)
	go f("goroutine")

	//anonymous function has been used to start a goroutine
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	var input string
	fmt.Scanln(&input)
	fmt.Println(input, "done")
}
