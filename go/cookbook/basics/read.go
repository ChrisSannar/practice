package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
* NOTE:
* - Always validate/sanitize user inputs
* - consider trimming (`strings.TrimSpace`)
* - Use buffer for large inputs
* - Provide feedback
* - Flush your buffers
* - Watch your input delimiters (`\n`)
* - `bufio.Reader` is better for performance
* - Large types? Pre-allocate buffer
* - Optimize with concurrency if needed
 */
func main() {
	fmt.Println("Read")
	// basic()
	// fmtInput()
	numberInput()
}

func basic() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Ender your name: ")

	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}
	name = strings.TrimSpace(name)
	fmt.Printf("Hello, %s", name)
}

func fmtInput() {
	var name string
	fmt.Print("(fmt) Enter your name: ")

	_, err := fmt.Scanln(&name)
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}
	fmt.Printf("Hello, %s\n", name)
}

func numberInput() {
	var input string
	fmt.Print("Enter your age: ")
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}

	age, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		fmt.Println("Invalid input, please enter a number")
		os.Exit(1)
	}

	fmt.Printf("Your age is %d.\n", age)
}
