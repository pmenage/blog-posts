package gcsbucket

import (
	"blog-posts/gcsbot/botservice/pubsub"
	"fmt"

	"github.com/shomali11/slacker"
	"golang.org/x/net/context"
)

// Handle handles the bot command
func Handle(request *slacker.Request, response slacker.ResponseWriter) {
	bucket := request.Param("bucket")
	channel := request.Event.Channel
	if bucket == "" {
		response.Reply("Usage: @configbot gcsbucket bucket")
		return
	}
	ctx := context.Background()
	pbs.SendPubsubMessage(ctx, bucket, channel, "CREATE_GCS_BUCKET")
	response.Reply(fmt.Sprintf("Creating GCS bucket %v", bucket))
}
