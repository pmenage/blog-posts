package main

import (
	"encoding/json"
	"os"
	"time"

	"cloud.google.com/go/pubsub"
	"cloud.google.com/go/storage"
	"golang.org/x/net/context"
)

var (
	projectID      = os.Getenv("PROJECT_ID")
	topicID        = os.Getenv("TOPIC_ID")
	subscriptionID = os.Getenv("SUBSCRIPTION_ID")
)

// Message is the Message structure
type Message struct {
	Bucket  string `json:"bucket"`
	Event   string `json:"event"`
	Channel string `json:"channel"`
}

// Receive gets messages from Pub/Sub
func Receive(ctx context.Context, startConfig func(Message)) {
	var msg Message
	pubsubcli, err := pubsub.NewClient(context.Background(), projectID)
	if err != nil {
		panic(err)
	}
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
	err = sub.Receive(ctx, func(ctx context.Context, m *pubsub.Message) {
		err := json.Unmarshal([]byte(m.Data), &msg)
		if err != nil {
			panic(err)
		}
		m.Ack()
		startConfig(msg)
	})
	if err != nil {
		panic(err)
	}
}

// PublishMessage publishes messages on Pub/Sub
func PublishMessage(ctx context.Context, msg Message) {
	pubsubclient, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		panic(err)
	}
	topic := pubsubclient.Topic(topicID)
	byt, err := json.Marshal(msg)
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

func main() {
	ctx := context.Background()
	for {
		cctx, cancel := context.WithTimeout(ctx, 60*time.Second)
		Receive(cctx, func(msg Message) {
			if msg.Event == "CREATE_GCS_BUCKET" {
				storageclient, _ := storage.NewClient(ctx)
				bucketHandle := storageclient.Bucket(msg.Bucket)
				err := bucketHandle.Create(ctx, projectID, nil)
				if err != nil {
					msg.Event = "GCS_BUCKET_CREATION_FAILED"
				} else {
					msg.Event = "GCS_BUCKET_CREATED"
				}
				PublishMessage(ctx, msg)
			}
		})
		cancel()
	}
}
