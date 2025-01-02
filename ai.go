package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type OpenAIChatRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OpenAIResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func NewOpenAIChatRequest(groceriesAsString string) *OpenAIChatRequest {
	return &OpenAIChatRequest{
		Model: "gpt-4o",
		Messages: []Message{
			{Role: "system", Content: "You are a helpful assistant specializing in grocery categorization. Prefix each item in a grocery list with a single emoji representing its store location. Remember the emoji should represent the category NOT the item. Sort items by their prefixed emoji. Do not include category headers. If an item's category is unclear, default to 🛒 pantry. Correct spelling errors and capitalize item names. The valid categories are: produce, dairy, pantry, frozen, meat, toiletries. For pantry use 🛒, for produce use 🍑, for meat use 🥩, for frozen use 🧊, for toiletries use 🧴"},
			{Role: "user", Content: groceriesAsString},
		},
	}
}

func (request *OpenAIChatRequest) Send() (string, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")

	if apiKey == "" {
		return "", fmt.Errorf("OPENAI_API_KEY environment variable not set")
	}

	payload, err := json.Marshal(request)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %v", err)
	}

	url := "https://api.openai.com/v1/chat/completions"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return "", fmt.Errorf("failed to create HTTP request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send HTTP request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("OpenAI API returned status: %s, body: %s", resp.Status, string(body))
	}

	var response OpenAIResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return "", fmt.Errorf("failed to unmarshal response: %v", err)
	}

	if len(response.Choices) == 0 || response.Choices[0].Message.Content == "" {
		return "", fmt.Errorf("no content returned from OpenAI")
	}

	return strings.TrimSpace(response.Choices[0].Message.Content), nil
}
