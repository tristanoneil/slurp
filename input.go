package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type GroceryCollector struct {
	Groceries []string
}

func NewGroceryCollector() *GroceryCollector {
	return &GroceryCollector{
		Groceries: []string{},
	}
}

func (gc *GroceryCollector) CollectInput() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		scanner.Scan()

		input := scanner.Text()

		if strings.ToLower(input) == "done" {
			break
		}

		items := strings.Split(input, ",")

		for _, item := range items {
			trimmed := strings.TrimSpace(item)

			if trimmed != "" {
				gc.Groceries = append(gc.Groceries, trimmed)
			}
		}
	}
}

func (gc *GroceryCollector) GroceriesAsString() string {
	return fmt.Sprintf("- %s", strings.Join(gc.Groceries, "\n- "))
}
