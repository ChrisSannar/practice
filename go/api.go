package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	// client, err := genai.NewClient(ctx, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	fmt.Println(ctx)

	// result, err := client.Models.GenerateContent(
	// 	ctx,
	// 	"gemini-3-flash-preview",
	// 	genai.Text("Explain your usage limits and how to implement my own model using an api key"),
	// 	nil,
	// )
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(result.Text())
}
