package main

import "fmt"

func main() {

	m := make(map[string]int)
	fmt.Println(m)
	m["k1"] = 7
	m["k2"] = 9

	fmt.Println(m, m["k1"], len(m))

	//delete using delete builtin

	delete(m, "k1")
	fmt.Println(m)

	// access deleted key
	mm, prs := m["k1"]
	fmt.Println(mm, prs)
}
