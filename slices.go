package main

import "fmt"

func main() {
	// slices are a more powerful interface to sequences than arrays

	// slices are typed only by the elements they contain and not their number
	// use the make builtin
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	fmt.Println("not so empty:", s)
	s[2] = "rr"
	fmt.Println("not so empty:", s)
	fmt.Println("len:", len(s))

	//slices support append, which returns a slice containing one or more new
	//values
	s = append(s, "d")
	ss := append(s, "e", "f")
	fmt.Println("s is", s)
	fmt.Println("ss is", ss)

	// slices can be copied as well.
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy", c)

	//slices support a slice operator like python
	cl1 := c[0:4]
	fmt.Println("slice", cl1)
	cl2 := c[:2]
	fmt.Println("slice", cl2)

	// slicing in place doesn't change the original slice data. onl changes
	// where the slice's ptr points to in the array
	// https://blog.golang.org/go-slices-usage-and-internals
	c = c[0:2]
	fmt.Println("sliced original: ", c)
	c = c[:cap(c)]
	//this shit cool as shit
	fmt.Println("restored original", c)
	/// single line slice declare and init
	t := []string{"g", "h", "i"}
	fmt.Println(t)

	//multi dimensional slices can have varying lengths in each inner dimension
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i*i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
