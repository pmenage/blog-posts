# Creating a GCS bucket with a Golang Slack bot

Files for the [blog post](https://medium.com/unacastlabs/writing-a-slack-chatbot-in-golang-31758cba86fe) on Medium called Creating GCS buckets with Slack bots in Go

## Requirements

- Have Go installed
- Have a Google Cloud Platform project ready to use with Pub/Sub and Google Cloud Storage

Environment variables:
- `API_TOKEN` for Slack Bot (Create an app [here](https://api.slack.com/apps))
- `PROJECT_ID` is your GCP project ID
- `TOPIC_ID` is the name of the topic you're going to send the messages to (create a topic before running the example)
- `SUBSCRIPTION_ID` is a subscription to a Pub/Sub topic

## Install and run

Open two terminals, open the `botservice` folder in one and the `creationservice` in the other. Set the environment variables in both terminals, for example: `export API_TOKEN=<YourSlackApiToken>`. 

Only the `SUBSCRIPTION_ID` has to be different or else the services will listen to the same subscription and only one of them will get the messages. You can set them to anything, the example will create them if they do not exist.

Run `go get -d ./...` at the root of the example to get the package dependencies.

Run `go run main.go` in each of the terminals. One will start the bot, which will start listening to Slack. The other will start listening to Pub/Sub, waiting for the bot service to ask it to create GCS buckets. Add your bot to the channels in which you wish to use it and test the commands.