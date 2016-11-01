package main

import "fmt"
import "sort"
import "strconv"

// Interfaces
// in Go, interfaces are sets of methods that we use to define a set of actions
// if a Student type implements:
// - SayHi(), Sing() and BorrowMoney() as its methods
// and an Employee type has 3 methods implemented:
// - SayHi(), Sing() and SpendSalary() as its methods

// Student and Employee implement the interface: SayHi() and Sing()
type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human  //anonymous field
	school string
	loan   int
}

type Employee struct {
	Human
	company string
	money   int
}

// Add a sayhi method to Human
func (h Human) SayHi() {
	fmt.Println("Hi, my name is ", h.name, " and my age is ", h.age)
}

// Overload the sayhi method from the Human struct for employee
func (e Employee) SayHi() {
	fmt.Println("Good Morning, I'm a pompous piece named ", e.name, "and I work at ", e.company)
}

func (h Human) Sing(lyrics string) {
	fmt.Println("This human can sing ", lyrics)
}

func (h Human) Guzzle(beers string) {
	fmt.Println("Imma ddrink this beer dawg ", beers)
}

func (s Student) BorrowMoney(amount int) {
	s.loan += amount
}

func (e *Employee) SpendMoney(amount int) {
	e.money -= amount
}

// Let's define our first interface now An interface is a set of abstract
// methods, and can be implemented by non-interface types. It cannot therefore
// implement itself.
type Men interface {
	SayHi()
	Sing(lyrics string)
	Guzzle(beers string)
}

type YoungChap interface {
	SayHi()
	Sing(song string)
	BorrowMoney(amoun int)
}

type OldGent interface {
	SayHi()
	Sing(song string)
	SpendMoney(amount int)
}

// fmt.Println can accept any type of argument. Any type that implements the
// Stringer interface can be passed to fmt.Println as an argument implement the
// Stringer interface for Human. The Stringer interface contains one method - String()
// Override Stringer's String() method for a pretty print of the Human object
func (h Human) String() string {
	return "Name: " + h.name + ", Age: " + strconv.Itoa(h.age) + ", phone: " + h.phone
}
func (p person) String() string {
	return "Name: " + p.name + ", Age: " + strconv.Itoa(p.age)
}

// Embedded Interfaces

// Just like structs can be used as anonymous fields in structs, we can have
// interfaces as anonymous fields in other interfaces. They are called embedded
// interfaces.
type Interface interface {
	sort.Interface
	Push(x interface{}) //a Push method to push elements into a heap
	Pop() interface{}   // a Pop mthod that pops elements from a heap
}

/// Empty interfaces are super useful for collections of different types
type Element interface{}
type List []Element

type person struct {
	name string
	age  int
}

func main() {

	mark := Human{"Mark", 22, "2245532244"}
	mark.SayHi()

	paul := Student{Human{"Paul", 26, "111-222-XXX"}, "Harvard", 100}
	sam := Employee{Human{"Sam", 36, "444-222-XXX"}, "Golang Inc.", 1000}
	Tom := Employee{Human{"Sam", 36, "444-222-XXX"}, "Things Ltd.", 5000}

	var itrfc Men // define an interface

	itrfc = paul
	paul.SayHi()
	fmt.Println("This is Paul. a student: ")
	itrfc.SayHi()
	itrfc.Guzzle("Smirnof")
	itrfc.Sing("I'm a sore losey phoosey loser")

	fmt.Println(sam, Tom)

	// We can use interfaces to store collections of types that implement that
	// interface. For example, the Men interface is implemented by Student and
	// Employee So let's 'make' a slice of Men

	m := make([]Men, 3)
	m[0], m[1], m[2] = paul, sam, Tom
	fmt.Println(m)

	for _, value := range m {
		value.SayHi()
	}

	// Empty Interfaces don't contain any methods, and thus, every type
	// satisfies that interface
	var emptyi interface{}
	var i int = 5
	s := "haha"
	// We can store any type in an empty interface
	emptyi = s
	fmt.Println(emptyi)
	emptyi = i
	fmt.Println(emptyi)

	Bob := Human{"Bob", 39, "404-323-6666"}
	fmt.Println("this human is ", Bob)

	// this is a slice that can accept any type because List type is made of
	// Element, which is an empty interface
	list := make(List, 3)
	list[0] = 1
	list[1] = "Hello"
	list[2] = person{"Dennis", 32}

	// Type checking in Go maybe done using the Assertion of Comma-ok pattern
	// value, ok := element.(T) where element is the interface variable, T is
	// the type of assertion, value is variable value and ok is a boolean that
	// holds the value of the assertion
	for index, elem := range list {
		if value, ok := elem.(int); ok {
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		} else if value, ok := elem.(string); ok {
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		} else {
			fmt.Printf("list[%d] is of an unknown type\n", index)
		}
	}

	//Type checking of an empty interface can also be done using switch
	for index, element := range list {
		switch value := element.(type) {
		case int:
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		case string:
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		case person:
			fmt.Printf("list[%d] is a person and their value is %s\n", index, value)
		default:
			fmt.Printf("list[%d] is of an unknown type\n", index)

		}
	}
	// element.(type) pattern can only be used inside a switch body. Otherwise
	// use the comma-ok pattern described above. value, ok = element.(T)
}
