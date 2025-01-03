package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
	"path/filepath"
)

type Config struct {
	Prompt string `toml:"prompt"`
}

func main() {
	fmt.Println("Welcome to Slurp! What are we shopping for this week? 🍜")
	fmt.Println("Add groceries line by line or comma separated. When you're finished type 'done' and <return>.")

	config, err := loadConfig()
	if err != nil {
		fmt.Printf("Error loading config %v\n", err)
		return
	}

	gc := NewGroceryCollector()
	gc.CollectInput()

	request := NewOpenAIChatRequest(config.Prompt, gc.GroceriesAsString())

	responseText, err := request.Send()
	if err != nil {
		fmt.Printf("Error communicating with OpenAI: %v\n", err)
		return
	}

	fmt.Println("\nCategorized Groceries:")
	fmt.Println(responseText)
}

func loadConfig() (Config, error) {
	configPath := filepath.Join(os.Getenv("HOME"), ".config", "slurp", "config")

	var config Config
	if _, err := toml.DecodeFile(configPath, &config); err != nil {
		return Config{}, err
	}

	return config, nil
}
