package main

import (
	"fmt"
)

/**
* NOTE:
* - Use type embedding to promote composition
* - Make sure types are related before embedding
* - Be cautious with colisions
* - type embedding assumes sharing methods
* - Use type embedding for method reuse
* - Use sensibly. Possible to have minor runtime costs
 */
func main() {
	fmt.Println("typeEmbed")
	// basics()
	// typeEmbedding()
	anonVsNamed()
}

type Animal struct {
	Name string
}

func (a *Animal) Speak() string {
	return "..."
}

type Dog struct {
	Animal
	Breed string
}

func basics() {
	dog := Dog{Animal: Animal{Name: "Bobby"}, Breed: "Husky"}
	fmt.Printf("%s is a %s and says: %s\n", dog.Name, dog.Breed, dog.Speak())
}

type Speaker interface {
	Speak() string
}

type Parrot struct{}

func (p Parrot) Speak() string {
	return "Hello! I'm a Parrot! Squak!"
}

type Human struct{}

func (h Human) Speak() string {
	return "Hi! I'm a Human!"
}

func saySomething(s Speaker) {
	fmt.Println(s.Speak())
}

func typeEmbedding() {
	parrot := Parrot{}
	human := Human{}

	saySomething(parrot)
	saySomething(human)
}

type Engine struct {
	Power int
}

type Car2 struct {
	Engine
	Make          string
	EngineDetails Engine
}

func anonVsNamed() {
	myCar := Car2{Engine: Engine{Power: 150}, Make: "Tesla", EngineDetails: Engine{Power: 200}}
	fmt.Printf("Car Make: %s, Power: %d, Detailed Power: %d\n", myCar.Make, myCar.Power, myCar.EngineDetails.Power)
}
