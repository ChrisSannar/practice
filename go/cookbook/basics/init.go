package main

import (
	"fmt"
)

// Runs when this package is accessed
func init() {
	fmt.Println("Package initialization")
}

/*
* NOTE:
* - Limit `init()` use. Better to be explicit in `main()`
* - always initialize else you'll get a `nil` poitner (and zero values)
* - nil map: read? ok, write? panic!
* - nil slice: range? ok, append? ok, index? panic!
* - If size is known, use literal syntax (`make(...)`)
 */
func main() {
	fmt.Println("Go Init")
	// basic()
	// initStruct()
	// initSlice()
	// initMap()
}

func basic() {
	count := 10
	name := "Alex"
	var level int = 80
	var country string = "USA"

	fmt.Println(count, name, level, country)
}

type User struct {
	Name string
	Age  int
}

func initStruct() {
	user1 := User{Name: "Alice", Age: 25}
	user2 := User{"Bob", 30}
	fmt.Println(user1, user2)
}

func initSlice() {
	slice1 := make([]int, 5)
	slice2 := []int{1, 2, 3, 4, 5}
	fmt.Println(slice1, slice2)
}

func initMap() {
	map1 := make(map[string]int)
	map1["key1"] = 1
	map2 := map[string]int{"key2": 2, "key3": 3}
	fmt.Println(map1, map2)
}
