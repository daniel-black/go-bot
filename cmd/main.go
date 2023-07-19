package main

import (
	"flag"
	"go-bot/pkg/app"
	c "go-bot/pkg/constants"
	"go-bot/pkg/openaiclient"
	"net/http"
)

// Ex: go run cmd/main.go -gpt -temp=0.8 -max=1000 -sm="You are a pro programmer. help me."
func main() {
	// Optional flags can specify model and system message
	isGpt4Ptr := flag.Bool("gpt4", false, c.IS_GPT_4_USAGE)
	tempPtr := flag.Float64("temp", c.DEFAULT_TEMP, c.TEMP_USAGE)
	maxPtr := flag.Int("max", c.DEFAULT_MAX, c.MAX_USAGE)
	systemMessagePtr := flag.String("sm", c.DEFAULT_SYSTEM_MESSAGE, c.SYSTEM_MESSAGE_USAGE)

	// Parse the flags
	flag.Parse()

	// Determine which model to use
	var model string
	if *isGpt4Ptr {
		model = c.GPT_4
	} else {
		model = c.GPT_3_5
	}

	// Create a new OpenAIClient
	openAIClient := &openaiclient.OpenAIClient{
		BaseURL:     c.API_URL,
		HTTPClient:  &http.Client{},
		APIKey:      c.API_Key,
		Model:       model,
		Temperature: *tempPtr,
		MaxTokens:   *maxPtr,
	}

	// Create a new App
	application := app.NewApp(openAIClient, *systemMessagePtr)

	// Run the app
	application.Run()
}
