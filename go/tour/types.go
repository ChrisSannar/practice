package main

import (
	"fmt"
	"strings"
	"math"
)

// import "golang.org/x/tour/pic"
import "golang.org/x/tour/wc"

type Vertex struct {
	X, Y int
}

func pointers() {
	fmt.Println("Pointers")
	i, j := 42, 2701

	// '&' creates a reference pointer of the value
	p := &i
	fmt.Printf("p: %T\n", p)	// The type here is *int
	fmt.Println("p: ", p)	// this points out th:e memory address
	fmt.Println("*p: ", *p) // this will print the value of 'i'
	*p = 21	// This then resets the value of I becuase the *p dereferences it
	fmt.Println("i: ", i)
	fmt.Println()

	p = &j	// Now p is a referene to j
	*p = *p / 37 // *p becomes -> j
	fmt.Println("j: ", j)
}

func structs() {
	fmt.Println("Structs")
	v := Vertex{1, 2}
	fmt.Println(v)
	v.X = 4
	fmt.Println(v.X)

	// Structs also operate similar with pointers
	p := &v
	p.Y = 1e9	// no dereferencing needed
	fmt.Println(v)

	v2 := &Vertex{X: 14}	// Y: 0 is implied
	fmt.Println("v2 ", v2)
}

// Arrays need the <name>[<size>]<type> context to be initialized
func arrays() {
	fmt.Println("Arrays: ")
	var strs[2]string
	strs[0] = "Hello"
	strs[1] = "World"
	fmt.Println(strs)

	fmt.Println()
	fmt.Println("Slices:")
	fib := [6]int{1, 1, 2, 3, 5, 8}
	fmt.Println(fib)

	// array slices are allowed as well
	sliced := fib[:2]	// slices up till the index
	fmt.Println(sliced)

	sliced[0] = -1	// Slices are references, so the ref values will be changed
	fmt.Println(fib)

	fmt.Println()
	// More examples of how slices manage the same way
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	fmt.Println()
	// Arrays can be assigned by literals as well
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	// A number size for the array isn't needed if all the contents are assigned then initialized
	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	// This can also be done with structs
	st := []struct{
		i int
		b bool
	}{
		{2, true}, 
		{3, false},
		{5, true}, 
		{7, true}, 
		{11, false},
		{13, true},
	}
	fmt.Println(st);

	fmt.Println()
	// Slice defaults are allowed as well
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)
	s = s[1:4]	// Droping any of the fist items shorten the capacity 
	printSlice(s)
	s = s[:2]	// Droping the end items doesn't affect capacity
	printSlice(s)
	s = s[1:]
	printSlice(s)
	// These would be all the same: s[0:5], s[:5], s[0:], s[:]
	
	var s_nil []int
	printSlice(s_nil)
	if (s_nil == nil) { fmt.Println("nil!") }

	fmt.Println()
	fmt.Println("Make: ")
	// Creating Slices with `make`
	am := make([]int, 5)
	printSlice(am)
	// args: (type, lenght, capacity)
	bm := make([]int, 0, 5)
	printSlice(bm)

	fmt.Println()
	fmt.Println("Multi-Arrays:")
	// Slices can include any time
	board := [][]string{
		[]string{"-", "-", "-"},
		[]string{"-", "-", "-"},
		[]string{"-", "-", "-"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][1] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
	fmt.Println()

	fmt.Println("Append: ")
	// You can append values to the array using `append`
	// Any append beyond capacity will double the capacity
	printSlice(append(am, 4))
	printSlice(append(bm, 2, 4, 6))
	fmt.Println()

	fmt.Println("Range: ")
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d == %d\n", i, v)
	}

	fmt.Println()
	pow2 := make([]int, 10)
	for i := range pow2 {
		pow2[i] = 1 << uint(i)
	}
	for _, val := range pow2 {
		fmt.Printf("%d\n", val)
	}

}

type Coords struct {
	Lat, Long float64
}

func maps() {
	fmt.Println("Maps: ")

	var m map[string]Coords // Default maps == nil
	m = make(map[string]Coords) // Need to be initialized with `make`
	m["Bell Labs"] = Coords{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
	
	// Maps can be initialized with literals
	var m2 = map[string]Coords{
		"Google": Coords{
			37.42204, -122.08408,	// <-- need that trailing comma
		}, // <-- This one too
	}
	fmt.Println(m2["Google"])

	// If the value is implied (Coords) then a literal can be used ({0.0, 0.0,})
	var m3 = map[string]Coords{
		"Nowhere": {0.0, 0.0,},
	}
	fmt.Println(m3["Nowhere"])

	m4 := make(map[string]int)
	m4["Answer"] = 42
	fmt.Println("Val: ", m4["Answer"])

	m4["Answer"] = 48
	fmt.Println("Val changed: ", m4["Answer"])
	
	delete(m4, "Answer")
	fmt.Println("Val: ", m4["Answer"])	// Deleted values default (0)

	//	`v` will be default if deleted, `ok` confirms (true/false)
	v, ok := m4["Answer"]
	fmt.Println("Val: ", v, "Present?", ok)	
	wc.Test(WordCount)
}

func functions() {
	fmt.Println("Functions: ")

	// Functions can be values as well
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	// Passed as variables into other functions
	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
	fmt.Println()

	fmt.Println("Closures: ")
	pos, neg := addr(), addr()	// Both pos, and neg share the `sum` value in the closeure because they're both `addr()`
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
	fmt.Println()

	test := closureTest()
	test()()
	fmt.Println()

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

func fibonacci() func() int {
	sum1, sum2 := 0
	return func() int {
		if sum1 == 0 && sum2 == 0 {
			sum1 = 1
			return 0
		}
		result := sum1 + sum2
		sum1 = sum2
		sum2 = result
		return result
	}
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// The funciton addr returns is a closure, having access to "sum"
func addr() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func closureTest() func() func() {
	test1 := "I'm test 1"
	func1 := func() func() {
		test2 := "I'm test 2"
		return func() {
			test3 := "I'm test 3"
			fmt.Println(test1, test2, test3)
		}
	}
	return func1
}

func printSlice(s []int){
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func Pic(dx, dy int) [][]uint8 {
	image := make([][]uint8, dy)
	for y := range image {
		image[y] = make([]uint8, dx)
		for x := range image[y] {
			image[y][x] = uint8(x ^ y)
		}
	}
	return image
}

func WordCount(s string) map[string]int {
	result := make(map[string]int)
	fields := strings.Fields(s)
	for _, val := range fields {
		_, inMap := result[val]
		if !inMap {
			result[val] = 1
		} else {
			result[val] += 1
		}
	}
	return result
}

func main() {
	/*
	pointers()
	fmt.Println()

	structs()
	fmt.Println()
	
	arrays()
	fmt.Println()

	// fmt.Println(Pic(2, 4))
	*/
	// pic.Show(Pic)
	// img := pic.Show(Pic) // This function takes a function of type (dx, dy int) [][]uint8
	// filter := img[5:]
	// fmt.Printf("%s", strings.Join(filter, ""))

	// maps()
	functions()
}
