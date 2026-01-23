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
}

func main() {
	fmt.Println("Methods and Interfaces:")
	methods()
}
