package main

import (
	"fmt"
	"image"
	"image/color"
	"io"
	"math"
	"os"
	"strings"
	"time"

	"golang.org/x/tour/pic"
	"golang.org/x/tour/reader"
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
	fmt.Println(v.Abs()) // Similar to classes, we can call the method on the struct
	fmt.Println(AbsFunc(v))

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	v.Scale(10) // It doesn't matter if the value called by the pointer reciever is a pointer or not
	fmt.Println("Scaled: ", v.Abs())
	ScaleFunc(&v, 5) // Since it's a normal funciton, we need to pass the vertex in by reference
	fmt.Println("Scaled Again: ", v.Abs())

	p := &v // The reverse is true for Methods, Go automatically interperates it as a pointer
	fmt.Println("Pointer v: ", p.Abs())

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

type Person struct {
	Name string
	Age  int
}

type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v",
		(ip[0]),
		(ip[1]),
		(ip[2]),
		(ip[3]))
}

type MyError struct {
	When time.Time
	What string
}

// Similar to `Stringer`, this is the message taken for errors
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

// Type `error` is the same as an `Exception` in java
func run(val interface{}) error {
	str, ok := val.(string)
	if ok {
		fmt.Printf("%v is a string!", str)
		return nil
	}
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func MySqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return math.Sqrt(x), nil
}

// The `String` method is the equivilant to `toString()` in java
// it's auto interperated via functions that take in string input (like fmt.Print)
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
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
	a = f  // `MyFloat` implements `Abser` with `func (f MyFloat) Abs() float64`
	a = &v // `*Vertex` implements `Abser` with `func (v Vertex) Abs() float64`

	a = v // `Vertex` does not implement `Abser`. It is just a struct.

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

	fmt.Println("Type Assertions:")
	var inter interface{} = "hello"
	s, ok := inter.(string) // You can asssert a type using this syntax + with a check
	fmt.Println(s, ok)      // ("hello", true)

	f2, ok := inter.(float64)
	fmt.Println(f2, ok) // (0, false)

	// This next line causes a panic since there is no check
	// f2 = i.(float64)
	fmt.Println(f2)

	doType(21)
	doType("Howdy")
	doType(true)
	fmt.Println()

	fmt.Println("Stringers: ")
	arthur := Person{"Arthur Dent", 42}
	zaphod := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(arthur, zaphod)

	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
	fmt.Println()

	fmt.Println("Errors: ")
	run("Stringy") // This will work because it has a string
	// but this wont because it's taking an int
	if err := run(42); err != nil {
		// Since error types are then handled accordingly
		fmt.Println(err)
	}

	fmt.Println(MySqrt(2))
	fmt.Println(MySqrt(-2))
	fmt.Println()

	fmt.Println("Readers: ")
	// Readers are basically streams
	r := strings.NewReader("Hello Reader!")
	b := make([]byte, 8) // This will act as a buffer
	for {
		// As the read the data in, `n` is how much data we read this time
		n, err := r.Read(b) // `err` will be `io.EOF` when finished
		fmt.Printf("n = %v\nerr = %v\nb= %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n\n", b[:n]) // Something to keep in mind: `b` still has the contents from the last read, it just writes over them to `n` index.
		if err == io.EOF {
			break
		}
	}

	reader.Validate(MyReader{})

	// rot13Reader: Rotation cypher by 13
	sRot := strings.NewReader("Lbh penpxrq gur pbqr!")
	rRot := rot13Reader{sRot} // The reader (feed of data) is then encapsulated by the rot13Reader
	io.Copy(os.Stdout, &rRot) // Spits it to standard output

	fmt.Println()
	fmt.Println()
}

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(b []byte) (n int, e error) {
	// First we're reading that section of the feed. Returns the index of the end
	msgBuffer, err := rot.r.Read(b)
	// For each item in that section of the buffer...
	for idx, val := range b[:msgBuffer] {
		// Ignore punctuation
		if val < 65 {
			continue
		}
		// Make sure we rotate by max ASCII value
		newVal := (val + 13) % 123
		// Lower bound of 'A'
		if newVal < 65 {
			newVal += 65
		}
		// Set the buffer data to the "rotated" value
		b[idx] = byte(newVal)
	}
	return len(b), err
}

type MyReader struct{}

func (myReader MyReader) Read(b []byte) (n int, e error) {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func describeEmpty(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// We can switch a type using this syntax
func doType(i interface{}) {
	switch v := i.(type) { // Note we have the use the keyword `type` here
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

type Image struct{}

func (img Image) Bounds() image.Rectangle {
	var rect image.Rectangle = image.Rect(0, 0, 100, 100)
	return rect
}
func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}
func (img Image) At(x, y int) color.Color {
	return color.White
}

func images() {
	fmt.Println("Images: ")

	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println("Bounds: ", m.Bounds())
	fmt.Print("At: ")
	fmt.Println(m.At(0, 0).RGBA())

	m2 := Image{}

	pic.ShowImage(m)
	fmt.Println()
	pic.ShowImage(m2)
	fmt.Println()
}

func main() {
	fmt.Println("Methods and Interfaces:")
	methods()
	interfaces()
	images()
}
