package main

import (
	"fmt"
	"log"
	"regexp"
)

import "github.com/go-playground/validator/v10"

/**
* NOTE:
* - Layer Validation: input at several layers in your application
* - Leverage 3rd party libraries (like validator used here)
* - Careful with Regexp as performance can suffer
* - Sanitize input
* - Always validate server-side
* - Consider edge cases: max length, unexpected symbols, binary characters, etc.
* - Make sure to handle errors properly
* - `regexp.MustCompile` is more performant
* - Minimal data copying
* - Concise errors: don't expose internal logic/sensitive info
 */
func main() {
	fmt.Println("Input")
	useValidate()
}

type User2 struct {
	Name  string `validate:"required"`
	Age   int    `validate:"gte=0,lte=130"`
	Email string `validate:"required,email"`
}

func useValidate() {
	v := validator.New()

	user := User2{
		Name:  "Alice",
		Age:   25,
		Email: "alice@example.com",
	}

	err := v.Struct(user)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			log.Fatal(err)
		}
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println("Validation Error:", err)
			log.Fatal(err)
		}
	} else {
		fmt.Println("User input is valid!")
	}
}

func useRegexp() {
	// email := "example@example.com"
	email := "asdf1234@1234.co.jp"
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if emailRegex.MatchString(email) {
		fmt.Println("Email address is valid.")
	} else {
		fmt.Println("Email address is invalid.")
	}
}
