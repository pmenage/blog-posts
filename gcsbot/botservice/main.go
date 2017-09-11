package main

import (
	"blog-posts/gcsbot/botservice/gcsbucket"
	"blog-posts/gcsbot/botservice/pubsub"

	"os"

	"golang.org/x/net/context"

	"github.com/shomali11/slacker"
)

func main() {
	slackAPIKey := os.Getenv("API_TOKEN")
	bot := slacker.NewClient(slackAPIKey)
	ctx := context.Background()

	go pbs.ReceiveAndSend(ctx, slackAPIKey)

	bot.Command("gcsbucket <bucket>", "Create a new GCS bucket", gcsbucket.Handle)
	bot.Listen()
}
