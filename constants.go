package main

import "fmt"
import "math"

//supportsconstants of char, str, bool and numerics
//use const to declare constant value
const s string = "constant"

func main() {

	fmt.Println(s)

	const n = 5000000000

	// arithmetic on consts ar eof arbitrary precision
	const d = 3e20 / n
	fmt.Println(d)
	//numeric const has no type until an explicit cast or till its given one
	fmt.Println(int64(d))

	// A number can be given a type by using it in a context that requires
	// one, such as a variable assignment or function call. For example,
	// here math.Sin expects a float64.
	fmt.Println(math.Sin(n))

}
