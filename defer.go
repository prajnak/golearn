package main

import "fmt"

// Defer is a go keyword. Statments prepended with defer get executed before the
// function returns. More than one defer statements will get executed in reverse order
// func ReadWrite() bool {
// 	file.Open("file")
// 	defer file.close()
// }
func main() {
	for i:=0; i<5; i++ {
		defer fmt.Println(i)
	}
}

