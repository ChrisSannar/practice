package main

import (
	"fmt"
	"reflect"
)

/**
* NOTE:
* - Use reflection sparingly
* - Always check if a value and type can be modifie
* - Always pass a pointer when you want to modify a value
* - Check for zero values (reflecting on `nil` panics)
* - Reflection is slower than basic operations
* - Cache frequent uses
 */
func main() {
	fmt.Println("Reflection")
	// basics()
	modify()
	fmt.Println()
	reflectStruct()
}

func basics() {
	var x float64 = 3.14
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Println("Type:", t)
	fmt.Println("Value:", v)
}

func modify() {
	var x float64 = 1
	v := reflect.ValueOf(&x).Elem()

	fmt.Println("Original Value:", v)

	v.SetFloat(2)
	fmt.Println("Modified Value:", v)
	fmt.Println("Underlying value:", x)
}

type Person2 struct {
	Name string
	Age  int
}

func reflectStruct() {
	p := Person2{"Alice", 25}
	val := reflect.ValueOf(&p).Elem()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fmt.Printf("Field %d: %v\n", i, field)
	}

	if nameField := val.FieldByName("Name"); nameField.IsValid() {
		if nameField.CanSet() {
			nameField.SetString("Bob")
		}
	}
}
