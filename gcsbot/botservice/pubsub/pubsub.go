package pbs

import (
	"encoding/json"
	"os"

	"cloud.google.com/go/pubsub"
	"github.com/nlopes/slack"
	"golang.org/x/net/context"
)

var (
	projectID      = os.Getenv("PROJECT_ID")
	topicID        = os.Getenv("TOPIC_ID")
	subscriptionID = os.Getenv("SUBSCRIPTION_ID")
)

// Message contains information in Pub/Sub message
type Message struct {
	Bucket  string `json:"bucket"`
	Event   string `json:"event"`
	Channel string `json:"channel"`
}

// SendPubsubMessage publishes a message on Pub/Sub
func SendPubsubMessage(ctx context.Context, bucket, channel, event string) {
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		panic(err)
	}
	topic := client.Topic(topicID)
	message := Message{
		Bucket:  bucket,
		Event:   event,
		Channel: channel,
	}
	byt, err := json.Marshal(message)
	if err != nil {
		panic(err)
	}
	_, err = topic.Publish(ctx, &pubsub.Message{
		Data: []byte(byt),
	}).Get(ctx)
	if err != nil {
		panic(err)
	}
}

// ReceiveAndSend receives then send a message
func ReceiveAndSend(ctx context.Context, slackAPIKey string) {
	for {
		pubsubcli, err := pubsub.NewClient(ctx, projectID)
		if err != nil {
			panic(err)
		}
		api := slack.New(slackAPIKey)
		var message Message
		topic := pubsubcli.Topic(topicID)
		sub := pubsubcli.Subscription(subscriptionID)
		exists, err := sub.Exists(ctx)
		if !exists {
			sub, err = pubsubcli.CreateSubscription(ctx, subscriptionID, pubsub.SubscriptionConfig{
				Topic: topic,
			})
			if err != nil {
				panic(err)
			}
		}

		cctx, cancel := context.WithCancel(ctx)
		sub.Receive(cctx, func(ctx context.Context, m *pubsub.Message) {
			err := json.Unmarshal([]byte(m.Data), &message)
			if err != nil {
				panic(err)
			}
			m.Ack()
			cancel()
		})
		switch message.Event {
		case "GCS_BUCKET_CREATED":
			api.PostMessage(string(message.Channel), "Bucket created for "+message.Bucket, slack.PostMessageParameters{})
		case "GCS_BUCKET_CREATION_FAILED":
			api.PostMessage(string(message.Channel), "Bucket not created for "+message.Bucket, slack.PostMessageParameters{})
		}
	}
}
