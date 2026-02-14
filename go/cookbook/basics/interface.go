package main

import (
	"fmt"
)

/**
* NOTE:
* - Use Names to reflect properties. Use verbs
* - Utilize assertion only when necessary (and handle gracefully)
* - Avoid large interfaces
* - Don't mix types in the same scope
* - Focus more on behavior, not concrete types
* - Interfaces have a pointer and a type, consider the data & memory usage
* - Consider avoiding polymorphism
 */
func main() {
	fmt.Println("Interfaces:")
	basics()
	typeAssert()
}

type Vehicle interface {
	Drive() string
}

type Car struct {
	Model string
}

func (c Car) Drive() string {
	return "All in!"
}

// We can type assert using *.(type)
func checkVehicle(v Vehicle) {
	if car, ok := v.(Car); ok {
		fmt.Printf("Car model: %s\n", car.Model)
	} else {
		fmt.Println("Unknown type")
	}
}

func typeAssert() {
	var v Vehicle
	v = Car{Model: "Renault"}
	checkVehicle(v)
}

type Greeter interface {
	Greet() string
}

type Human struct {
	Name string
}

func (h Human) Greet() string {
	return "Hello! " + h.Name
}

type Cow struct {
	Name string
}

func (c Cow) Greet() string {
	return "Moo? " + c.Name
}

// As long as the functions and properties match, it works
// Aka: "if it quacks like a duck..."
func basics() {
	var g Greeter
	human := Human{Name: "Alex"}
	cow := Cow{Name: "Milka"}

	g = human
	fmt.Println(g.Greet())

	g = cow
	fmt.Println(g.Greet())

}
