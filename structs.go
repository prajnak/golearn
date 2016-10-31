package main

import "fmt"

// structures are typed collections of fields They're useful for grouping data
// together to form records

// the persons struct has two fields. Each field has a name and type
type person struct {
	name string
	age  int
}

// Embedded fields in a struct have no names but have types
type Human struct {
	name string
	age int
	weight int
}

// this struct uses an embedded field
// this is a form of inheritance 
type Student struct {
	Human // this means the student struct will have all the fields that HUman
	// struct has
	specialty string
	Skills
	int //built in type as an embedded field
}

type Skills []string

func main() {

	// syntax to create a new struct --> person{"Bob", 20}

	// yield a pointer to the newly created struct
	pp := &person{"Bob", 20}

	ppp := person{"Bob", 20}
	fmt.Println(pp.age, ppp.age)

	// structs are mutable
	pp.age = 50
	fmt.Println(pp.age, ppp.age)

	var andy person
	andy.name = "bob"
	andy.age = 45
	fmt.Println(andy)

	mark := Student{Human: Human{"Mark", 22, 190}, specialty:"Arts", Skills:[]string{"dope", "damn"}}
	fmt.Println(mark, mark.name, mark.Human.age, mark.Skills)
	mark.int = 3
	fmt.Println(mark.int)
}

// When a struct embeds another struct in a field and both structs have a field
// with the same name, Outer fields get upper level access. ie., the outside
// struct's field will be accessible with that name. field overloading. The
// embedded struct's masked field maybe accessed explicitly with
// struct.EmbeddedField.Field construct
