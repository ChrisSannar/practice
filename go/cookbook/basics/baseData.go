package main

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	fmt.Println("baseData")
	fmt.Println()

	basicEncoding()
}

// NOTE Profile and benchmark different coding strategies to see which is best
func basicEncoding() {
	// Basic Encoding
	data := "Text to encode"
	encodedData := base64.StdEncoding.EncodeToString(([]byte(data)))
	fmt.Println("Encoded: ", encodedData)

	// Basic Decoding
	decodedData, err := base64.StdEncoding.DecodeString(encodedData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Decoded:", string(decodedData))

	// Basic JSON
	person := map[string]any{
		"name": "Jack",
		"age":  25,
	}
	jsonData, err := json.Marshal(person)
	if err != nil {
		log.Fatal()
	}
	fmt.Println("JSON Encoded", string(jsonData))
	var person2 map[string]any
	err2 := json.Unmarshal([]byte(jsonData), &person2)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Printf("JSON Decoded: %+v\n", person)
	// NOTE: For larger JSON sets, use `json.Decoder` for better performance

	// Hex Encoding
	hexData := "Text to encode to hex"
	encodedHex := hex.EncodeToString([]byte(hexData))
	fmt.Println("Hex Encoded: ", encodedHex)
	decodedHex, err3 := hex.DecodeString(encodedHex)
	if err3 != nil {
		log.Fatal(err3)
	}
	fmt.Println("Hex Decoded:", string(decodedHex))
}
