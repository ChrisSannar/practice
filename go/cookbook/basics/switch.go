package main

import (
	"fmt"
)

/**
* NOTE:
* - Use fallthrough sparingly
* - Switch without expression defaults to `switch true`
* - Make sure to handle all cases possible
* - Opt for `switch` over `if-else` for multiple values
* - Put most likely cases first for performanc
* - Avoid complex expressions (evaluate before)
 */
func main() {
	fmt.Println("Switch")
	// basic()
	// conditions()
	// multiple()
	fallThrough()
}

func basic() {
	dayOfWeek := 3

	switch dayOfWeek {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	case 8:
		fmt.Println("Invalid day of week")
	}
}

func conditions() {
	score := 85

	switch {
	case score >= 90:
		fmt.Println("Grade: A")
	case score >= 80:
		fmt.Println("Grade: B")
	case score >= 70:
		fmt.Println("Grade: C")
	case score >= 60:
		fmt.Println("Grade: D")
	default:
		fmt.Println("Grade: F")
	}
}

func multiple() {
	day := "Saturday"

	switch day {
	case "Saturday", "Sunday":
		fmt.Println("I's the weekend!")
	default:
		fmt.Println("Weekday")
	}
}

func fallThrough() {
	age := 18

	switch {
	case age >= 18:
		fmt.Println("Adult")
		fallthrough // Golang switches break, unless there is this
	case age >= 13:
		fmt.Println("Teenager")
	default:
		fmt.Println("Child")
	}
}
