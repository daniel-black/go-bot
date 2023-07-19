package main

import (
	"go-bot/pkg/app"
	"go-bot/pkg/openaiclient"
)

const (
	// Put your OpenAI key here
	apiKey = ""
	apiUrl = "https://api.openai.com/v1/chat/completions"
	model  = "gpt-3.5-turbo"
)

// Entry point of the application. Initializes app and starts main loop.
func main() {
	// Create a new OpenAIClient
	openAIClient := openaiclient.NewOpenAIClient(apiUrl, apiKey, model)

	// Create a new App
	application := app.NewApp(openAIClient)

	// Run the app
	application.Run()
}
