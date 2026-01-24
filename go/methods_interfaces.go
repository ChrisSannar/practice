package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

type MyFloat float64

// Methods take a "reciever" (in our case, `(v Vertex)`) and attach the function to it.
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Methods are still functions. Here is function implementation of Abs
func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Methods can also take pointer recievers so they can change the pointer contents
// Without the pointer `*`, `v` wouldn't change and remain the same
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// This is how the funciton would look normally
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Methods can be applied to types from the same package
// They can also have the same name as other functions because they're tied to a different reciever
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func methods() {
	fmt.Println("Methods: ")
	v := Vertex{3, 4}
	fmt.Println(v.Abs())	// Similar to classes, we can call the method on the struct
	fmt.Println(AbsFunc(v))
	
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
	
	v.Scale(10)	// It doesn't matter if the value called by the pointer reciever is a pointer or not
	fmt.Println("Scaled: ", v.Abs())
	ScaleFunc(&v, 5)	// Since it's a normal funciton, we need to pass the vertex in by reference 
	fmt.Println("Scaled Again: ", v.Abs())

	p := &v	// The reverse is true for Methods, Go automatically interperates it as a pointer
	fmt.Println("Pointer v: ", p.Abs());

	// Notes: The 2 cases to use Pointer Recievers:
	//  1. We want to modify the original value passed into the method
	//  2. We don't want to copy a new value each time we call the function (as normal functions do)
	fmt.Println()
}

type Abser interface {
	Abs() float64
}

// It seems the philosophy of Go is to use the module itself to keep context
type I interface {
	M()
}

type T struct {
	S string
}

// This function then implements the interface `I` to `T`
func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println("M T:", t.S)
}

func (f MyFloat) M() {
	fmt.Println("M MyFloat:", f)
}

func interfaces() {
	fmt.Println("Interfaces: ")

	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	// In order to implement an interface, the given properties need to be applied
	// aka, interfaces are implemented implicitly (quacks like a duck...)
	a = f		// `MyFloat` implements `Abser` with `func (f MyFloat) Abs() float64`
	a = &v	// `*Vertex` implements `Abser` with `func (v Vertex) Abs() float64`

	a = v		// `Vertex` does not implement `Abser`. It is just a struct.

	fmt.Println(a.Abs())

	var i I = &T{"Message"}
	describe(i)
	i.M()

	// Now that i is of type `MyFloat`, it then sets it values to that specific types functions (like `M()`) 
	i = MyFloat(math.Pi)	
	describe(i)
	i.M()

	// Because t is nil, methods will still be called, just with a nil reciever
	// Go doesn't call a nullptr exception, instead it's up to methods to handle the exception
	var t *T
	i = t
	describe(t)
	i.M()

	var i2 I
	describe(i2)
	// nil interface value method calls result in runtime errors
	// i2.M() // <-- This would cause an error

	// `interface{}` is the Golang "any" type
	var iEmpty interface{}
	describe(i)

	iEmpty = 42
	describeEmpty(iEmpty)

	iEmpty = "Howdy"
	describeEmpty(iEmpty)

	fmt.Println()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func describeEmpty(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	fmt.Println("Methods and Interfaces:")
	methods()
	interfaces()
}
