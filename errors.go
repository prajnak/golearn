package main

import "fmt"
import "errors"
// go has an error type for dealing with error messages
// Go doesn't have a try-catch structure like Java
// It uses panic and recover to deal with errors.

// Panic is a built in function to break the normal flow of programs and get
// into panic status. When a function calls panic, its statements will stop
// being executed except for defer calls already registered by them

func main() {
	err := errors.New("emit this message that's hopefully helpful")
	if err != nil {
		fmt.Print(err)
	}
}
