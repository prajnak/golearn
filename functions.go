package main

import "fmt"

func plus(a int, b int) int {

	return a + b
}

func plusPlus(a, b, c int) int {

	return a + b + c
}

//multi return values
func vals() (int, int) {

	return 3, 7
}

// variadic functions
// can have any number of trailing arguments
func sum(nums ...int) {

	fmt.Println(nums)
	fmt.Println(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

// closures
func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

// Functions as values and types. Functions can be defined with the type keyword
// type typeName func(input1 inputType1, input2 inputType2[, ...]) (result1 resultType1 [, ...])
// lets us pass functions as values(??)
type testInt func(int) bool //define a function type variable
func isOdd(num int) bool {
	if num%2 == 0 {
		return false
	}
	return true
}

func isEven(num int) bool {
	if num%2 == 0 {
return true
	}
	return false
}

func filter(slice []int, f testInt) []int {

	var result []int // slice to store results
	for _,value := range slice {
		if f(value) {
			result = append(result, value) //slices are mutable
		}
	}
	return result
}


//recursion
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {

	res := plus(5, 4)
	fmt.Println(res)

	res = plusPlus(5, 4, 5)
	fmt.Println(res)
	a, b := vals()
	_, c := vals()

	fmt.Println(a, b, c)
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 4, 5, 2, 7, 8, 10000}
	// If you already have multiple args in a slice, apply them to a variadic
	// function using func(slice...) like this.
	sum(nums...)

	aa := intSeq()
	fmt.Println(aa())
	fmt.Println(aa())
	fmt.Println(aa())
	fmt.Println(aa())
	fmt.Println(aa())
	fmt.Println(aa())
	aa = intSeq()
	fmt.Println(aa())
	fmt.Println(aa())
	fmt.Println(aa())
	fmt.Println(aa())

	fmt.Println(fact(7))


	slc := []int {1,2,3,4,5}
	fmt.Println(slc)
	ab := filter(slc, isOdd)
	ac := filter(slc, isEven)
	fmt.Println(ab, ac)
}
