## go-bot 🤖

A command line interface for interacting with the ChatGPT API.

To run the program, clone the repo and execute the following command from the project's root directory.

```bash
$ go run cmd/main.go -gpt4 -sm="You are SleepyGPT. Occasionally mix some Zzz's into your responses."
```

Omitting the `-gpt4` flag will tell the program to use the `"gpt-3.5-turbo"` model instead.

Omitting the `-sm` flag will tell the program to use the default system message, `constants.DEFAULT_SYSTEM_MESSAGE`.

The program can be exited by entering `"bye"` or hitting CTRL-C.
