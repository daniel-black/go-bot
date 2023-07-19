package main

import (
	"flag"
	"go-bot/pkg/app"
	"go-bot/pkg/constants"
	"go-bot/pkg/openaiclient"
)

func main() {
	// Optional flags can specify model and system message
	isGpt4Ptr := flag.Bool("gpt4", false, "Is the model using gpt-4? If not, it will use gpt-3.5-turbo.")
	systemMessagePtr := flag.String("sm", constants.DEFAULT_SYSTEM_MESSAGE, "A system message to seed the AI with.")

	// Parse the flags
	flag.Parse()

	// Determine which model to use
	var model string
	if *isGpt4Ptr {
		model = constants.GPT_4
	} else {
		model = constants.GPT_3_5
	}

	// Create a new OpenAIClient
	openAIClient := openaiclient.NewOpenAIClient(constants.API_URL, constants.API_Key, model)

	// Create a new App
	application := app.NewApp(openAIClient, *systemMessagePtr)

	// Run the app
	application.Run()
}
