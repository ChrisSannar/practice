package main

import (
	"fmt"
)

/**
* NOTE:
* - Use pointers for large structs or when function need to modify data (especially for loops)
* - Always check pointer validity (test for nil)
* - Dereferencing a nil pointer results in a panic
* - Avoid excessive use. It gets complicated
* - Sometime addresses change, especially as programs scale
* - Dont pass pointers beyond their valid scope (causes memory leaks)
* - Investigate `sync.Pool`?
 */
func main() {
	fmt.Println("Pointers")
	// basics()
	// structPointer()
	funcArgs()
}

func basics() {
	var x int = 1
	var p *int = &x

	fmt.Println("Value of x:", x)
	fmt.Println("Address of x:", p)
	fmt.Println("Value via pointer p:", *p)

	*p = 2
	fmt.Println("New value of x:", x)
}

type Person struct {
	name string
	age  int
}

func structPointer() {
	p := Person{"Alice", 20}
	ptr := &p

	fmt.Println("Name:", p.name)
	fmt.Println("Age:", ptr.age)

	ptr.age = 25
	fmt.Println("Updated Age:", ptr.age)
}

func updateValue(num *int) {
	*num = 2
}

func funcArgs() {
	val := 1
	fmt.Println("Initial Value:", val)
	updateValue(&val)
	fmt.Println("Updated Value:", val)
}
