package main

import (
	"os"

	"github.com/shomali11/slacker"
)

func handle(request *slacker.Request, response slacker.ResponseWriter) {
	name := request.Param("name")
	if name == "" {
		response.Reply("Usage: @hellobot hello Name")
		return
	}
	response.Reply("Hey " + name + "!")
}

func main() {
	bot := slacker.NewClient(os.Getenv("API_TOKEN"))
	bot.Command("hello <name>", "Say hello to someone", handle)
	err := bot.Listen()
	if err != nil {
		panic(err)
	}
}
