# Slackbot

Files for the [blog post](https://medium.com/unacastlabs/writing-a-slack-chatbot-in-golang-31758cba86fe) on Medium called Writing a Slack ChatBot in Golang

## Requirements

- Have Go installed
- API_TOKEN for Slack Bot (Create an app [here](https://api.slack.com/apps))

## Install and run

Set the API token in an environment variable: `export API_TOKEN=<YourSlackApiToken>`

Clone one of the examples, and you can run `go get -d ./...` to get the package dependencies.

Run `go run main.go` to start your bot. Add your bot to the channels in which you wish to use it and test the commands.