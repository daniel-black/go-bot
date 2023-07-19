// The main application logic (main loop and handling of user input)
package app

import (
	"bufio"
	"encoding/json"
	"fmt"
	"go-bot/pkg/models"
	"go-bot/pkg/openaiclient"
	"io/ioutil"
	"net/http"
	"os"
)

type App struct {
	Client *openaiclient.OpenAIClient
}

func NewApp(client *openaiclient.OpenAIClient) *App {
	return &App{
		Client: client,
	}
}

func (app *App) Run() {
	reader := bufio.NewReader(os.Stdin)
	messages := make([]models.Message, 0, 10)

	for {
		// take user input
		fmt.Print("Me: ")
		userMessageContent, err := reader.ReadString('\n')
		if err != nil {
			os.Exit(1)
		}

		if userMessageContent == "bye" {
			fmt.Println("Goodbye!")
			os.Exit(0)
		}

		fmt.Print("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")

		// construct the user message
		userMessage := models.Message{
			Role:    "user",
			Content: userMessageContent,
		}

		// append it to the slice of messages
		messages = append(messages, userMessage)

		// make the request
		resp, err := app.Client.MakeChatRequest(messages)
		if err != nil {
			fmt.Printf("Error making request: %v\n", err)
			continue
		}

		// check the http status code
		if resp.StatusCode != http.StatusOK {
			fmt.Printf("Received non-OK HTTP status: %s\n", resp.Status)
			continue
		}

		// read the response body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error reading the response body: %v\n", err)
			continue
		}

		// unmarshal the response body into a ChatResponse instance
		var chatResponse models.ChatResponse
		err = json.Unmarshal(body, &chatResponse)
		if err != nil {
			fmt.Printf("Error unmarshalling response body: %v\n", err)
			continue
		}

		// close response body
		resp.Body.Close()

		// get the chat response
		aiMessage := chatResponse.Choices[0].Message

		// print the AI response
		fmt.Printf("AI: %s\n", aiMessage.Content)
		fmt.Print("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n")

		// append new message to messages slice
		messages = append(messages, aiMessage)

		// print whole response just for demo purposes
		// fmt.Printf("Received response :%+v\n", chatResponse)
	}
}
