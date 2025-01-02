package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Slurp! What are we shopping for this week? 🍜")
	fmt.Println("Add groceries line by line or comma separated. When you're finished type 'done' and <return>.")

	gc := NewGroceryCollector()
	gc.CollectInput()

	fmt.Println("You've added:")

	for _, item := range gc.Groceries {
		fmt.Println(item)
	}
}
