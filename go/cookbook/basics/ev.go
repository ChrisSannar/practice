package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("EV")

	// getEnv()
	// setEnv()
	lookupEnv()
}

func getEnv() {
	host := os.Getenv("TEMP")
	if host == "" {
		fmt.Println("TEMP is not set")
		return
	}
	fmt.Println("TEMP: ", host)
}

func setEnv() {
	err := os.Setenv("TEMP", "I'm temp2")
	if err != nil {
		log.Fatal(err)
	}

	appEnv := os.Getenv("TEMP")
	fmt.Println("CHANGED TEMP:", appEnv)
}

// NOTE: Lookup checks if it's unset, and not just an empty string
func lookupEnv() {
	if dbPort, exists := os.LookupEnv("DB_PORT"); exists {
		fmt.Println("Database port:", dbPort)
	} else {
		fmt.Println("DB_PORT is not set")
	}
}
