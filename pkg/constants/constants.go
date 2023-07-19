package constants

const (
	// Put your OpenAI key here
	API_Key                = ""
	API_URL                = "https://api.openai.com/v1/chat/completions"
	DEFAULT_SYSTEM_MESSAGE = "You are a knowledgeable and helpful assistant."
	GPT_3_5                = "gpt-3.5-turbo"
	GPT_4                  = "gpt-4"
	IS_GPT_4_USAGE         = "Is the model using gpt-4? If not, it will use gpt-3.5-turbo."
	TEMP_USAGE             = "A value between 0 and 1. Higher values like 0.8 will make the output more random, while lower values like 0.2 will make it more focused and deterministic."
	MAX_USAGE              = "The maximum number of tokens."
	SYSTEM_MESSAGE_USAGE   = "A system message to seed the AI with."

	DEFAULT_TEMP float64 = 0.8
	DEFAULT_MAX  int     = 500
)
