package main

import (
	"fmt"
	"log"
	"net/http"
)

/**
* NOTE:
* - Only use when needed
* - Log Panic
* - Graceful Shutdown: clear resources, network connections, use defer
* - Separate Error handling
* - Limit scope
* - Properly profile (benchmark, clear logs, analyze performance)
 */
func main() {
	fmt.Println("Panic!")
	// basic()
	httpPanic()
}

func httpPanic() {
	http.HandleFunc("/", safeHandler(exampleHandler))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func exampleHandler(w http.ResponseWriter, r *http.Request) {
	panic("Test panic")
}

func safeHandler(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Fprintf(w, "Recoverd from panic: %v", err)
			}
		}()
		h(w, r)
	}
}

func basic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from the panic:", r)
		}
	}()

	fmt.Println("Starting the program")
	causePanic()
	fmt.Println("This line won't be reached because of the panic")
}

func causePanic() {
	panic("Unable to proceed")
}
