package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Slurp! What are we shopping for this week? 🍜")
	fmt.Println("Add groceries line by line or comma separated. When you're finished type 'done' and <return>.")

	gc := NewGroceryCollector()
	gc.CollectInput()

	request := NewOpenAIChatRequest(gc.GroceriesAsString())

	responseText, err := request.Send()

	if err != nil {
		fmt.Printf("Error communicating with OpenAI: %v\n", err)
		return
	}

	fmt.Println("\nCategorized Groceries:")
	fmt.Println(responseText)
}
