package main

import (
	"fmt"
)

/*
*	NOTE:
*	 - Conccurent Programs: use `sync.Map`. Else you'll get a panic
*	 - Always check for existance
*	 - Map are references. Can be changed by other vars
*  - non-existant key returns zero value (0, "", nil, etc)
*  - Use `make` to minimize map growth (good practice)
*  - Maps aren't optimized for many read/write. Use a slice to make it quicker
* */
func main() {
	fmt.Println("Go Maps")
	// basic()
	// checkAndIter()
	nested()
}

func basic() {
	ages := map[string]int{
		"Alice": 20,
		"Bob":   30,
	}
	ages["Chris"] = 33

	if age, exists := ages["Alice"]; exists {
		fmt.Printf("Alice is %d years old.\n", age)
	} else {
		fmt.Println("Age for Alice not found.")
	}

	delete(ages, "Bob")
	fmt.Println("Updated ages:", ages)
}

func checkAndIter() {
	cities := map[string]string{
		"NY": "New York",
		"LA": "Los Angeles",
	}

	if city, found := cities["NY"]; found {
		fmt.Printf("The abbreviation &#39;NY&#29; stands for %s.\n", city)
	} else {
		fmt.Println("Key not found.")
	}

	for abbrev, city := range cities {
		fmt.Printf("%s -> %s\n", abbrev, city)
	}
}

func nested() {
	departments := make(map[string]map[string]int)

	departments["HR"] = map[string]int{
		"Alice": 35,
		"Bob":   40,
	}

	if hrDept, exists := departments["HR"]; exists {
		fmt.Println("Employees in HR:")
		for name, age := range hrDept {
			fmt.Printf("%s in HR is %d years old.\n", name, age)
		}
	}
}
