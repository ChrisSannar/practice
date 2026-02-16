package main

import (
	"encoding/json"
	"fmt"
	"log"
)

/**
* NOTE:
* - `omitempty` to avoid serializing empty fields
* - Clearly name JSON tags to ensure consistent API
* - Don't forget to export struct fields
* - Check your spelling
* - Make sure to handle `nil` values
* - Avoid deep loops which can consume significant CPU resources
 */
func main() {
	fmt.Println("annotations")
	basics()
}

type Person3 struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email,omitempty"`
}

func basics() {
	data := `{"name":"Alice", "age":25}`
	var p Person3
	if err := json.Unmarshal([]byte(data), &p); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Person: %+v\n", p)

	jsonData, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonData))
}
