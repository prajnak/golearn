package main

import "fmt"
import "reflect"

type rect struct {
	width, height int
}

// methods defined on struct types define a method with a receiver type the
// receiver type is a pointer to a rect struct
// Methods are functions with implicit first arguments as a type receiver
// Syntax: func (r ReceiverType) funcName(parameters) (results type) {}
func (r *rect) area() int {
	return r.width * r.height
}

// method receiver types can also be values instead of pointer. Fields changed
// inside the function aren't reflected in the original struct as this is
// passing the struct. Use a pointer receiver type if you need to change the
// original struct
func (r rect) perimeter() int {
	return r.width*2 + r.height*2
}

// Interfaces: they are a set of methods as well as a type. Let's create a new
// interface type called Animal.
type Animal interface {
	Speak() string
}

// Animal is any TYPE that has a Speak() method implemented on it A TYPE's
// implementation(satisfaction) of an interface is determined by Go
// automatically. There's no "implements" keyword in Go
// lets create a dog type that conforms to the Animal interface
type Dog struct {
	// this is empty
}

// create a method called Speak() for the Dog struct. It should use the Dog struct as a
// receiver type - either by type or value
func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {
}

func (c *Cat) Speak() string {
	return "Meow!"
}

// Method Inheritance
type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
}

type Employee struct {
	Human
	company string
}

func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s and you can bling me at %s", h.name, h.phone)
}

func main() {

	r := rect{width: 8, height: 10}
	fmt.Println("area: ", r.area())
	fmt.Println("perimeter: ", r.perimeter())

	// automatic conversion between values and pointers for method calls

	rp := &r
	fmt.Println(rp)
	fmt.Println(r)

	// print the type of r as well as rp
	fmt.Println(reflect.TypeOf(rp), reflect.TypeOf(r))

	mark := Student{Human{"Mark", 20, "22211122"}, "WCI"}
	andy := Employee{Human{"Andy", 40, "5533322"}, "WCI"}
	fmt.Println(mark)
	fmt.Println(andy)

	mark.SayHi() //even though the receiver type for sayhi is Human, Student,
	//which embeds a Human field can use the embedded struct's methods
}
