package main

import (
	"fmt"
	"log"
	"os"
)

// Even if main panics, defer will still try to run.
func main() {
	fmt.Println("Defer")

	// Use `defer` for operations to run at the end of the function
	f, err := os.Create("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = f.WriteString("File content")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File written successfully")

	// Multiple calls: runs FIFO
	defer fmt.Println("Bob")
	defer fmt.Println("Alice")
	fmt.Println("I'm After Bob and Alice")

	process()
}

func cleanup() {
	fmt.Println("Cleaning up resources")
}
func process() {

	// Defering functions
	defer cleanup()

	fmt.Println("Processing...")
	// Some kind of processing
	for i := 0; i < 5000000000; i++ {
		// NOTE: Go has little overhead with defer, but if it's tight, do cleanup inside
	}

}
