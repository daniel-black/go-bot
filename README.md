## go-bot 🤖

A command line interface for interacting with the ChatGPT API.

### To run the program

Clone the repo and execute the following command from the project's root directory.

```bash
$ go run cmd/main.go
```

The program can be exited by entering `"bye"` or hitting CTRL-C.

### Flags

##### `-gpt4`

Tells program to use `"gpt-4"` model instead of `"gpt-3.5-turbo"`

##### `-sm="A really useful system message of your choosing"`

A system message to seed the system. If omitted, a default system message (`constants.DEFAULT_SYSTEM_MESSAGE`) is used.

##### `-temp=0.8`

A value between 0 and 1. Higher values like 0.8 will make the output more random, while lower values like 0.2 will make it more focused and deterministic. Defaults to `0.8` if a value is not supplied.

##### `-max=500`

Sets the maximum number of tokens the API should respond with. Defaults to `1000` if a value is not supplied.

---

###### Example

A fully flagged-up command might look like:

```bash
$ go run cmd/main.go -gpt4 -sm="You are really smart." -temp=0.9 -max=3000
```
