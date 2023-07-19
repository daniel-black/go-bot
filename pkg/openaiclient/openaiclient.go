// Implements the HTTP request logic
package openaiclient

import (
	"bytes"
	"encoding/json"
	"go-bot/pkg/models"
	"net/http"
)

type OpenAIClient struct {
	BaseURL     string
	HTTPClient  *http.Client
	APIKey      string
	Model       string
	Temperature float64
	MaxTokens   int
}

func NewOpenAIClient(baseURL, apiKey, model string, temp float64, maxTokens int) *OpenAIClient {
	return &OpenAIClient{
		BaseURL:     baseURL,
		HTTPClient:  &http.Client{},
		APIKey:      apiKey,
		Model:       model,
		Temperature: temp,
		MaxTokens:   maxTokens,
	}
}

func (c *OpenAIClient) MakeChatRequest(messages []models.Message) (*http.Response, error) {
	// Create request body object
	body := models.ChatRequest{
		Model:       c.Model,
		Messages:    messages,
		Temperature: c.Temperature,
		MaxTokens:   c.MaxTokens,
	}

	// Marshal request body into JSON
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	// Create request
	req, err := http.NewRequest("POST", c.BaseURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.APIKey)

	// Execute the request
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
