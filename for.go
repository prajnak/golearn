package main

import "fmt"

func main() {

	// for is go's only ooping construct. 3 basic types of for loops
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// infinite loop
	for {
		fmt.Println("loop")
		break
	}
}
