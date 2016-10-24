package main

import "fmt"

func main() {

	// arrays are numbered sequence of elements of a specific length
	// a holds 5 ints. the array type includes element type and length
	// arrays are also zero valued if not initialized
	var a [5]int
	fmt.Println("empty:", a)

	// set a value at an index using array[index] = value. get is similar as
	// well

	a[4] = 100
	fmt.Println("set: ", a)
	fmt.Println("get: ", a[4])

	//length of array using the len function
	fmt.Println(len(a))

	//declare, init in single line like this
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("declared: ", b)

	// multi dimensional arrays are composable like this
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
