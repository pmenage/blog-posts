package main

import (
	"os"

	"github.com/shomali11/slacker"
)

func handle(request *slacker.Request, response slacker.ResponseWriter) {
	response.Reply("Hey!")
}

func main() {
	bot := slacker.NewClient(os.Getenv("API_TOKEN"))
	bot.Command("hello", "Say hello", handle)
	err := bot.Listen()
	if err != nil {
		panic(err)
	}
}
