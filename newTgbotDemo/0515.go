package main

import "github.com/NicoNex/echotron/v3"

const token = "6892958172:AAFCFpZ8RMI4MYpk6E8Vdpiubdly4kyYGtA"

func main() {
	api := echotron.NewAPI(token)

	for u := range echotron.PollingUpdates(token) {
		if u.Message.Text == "/start" {
			api.SendMessage("/string", u.ChatID(), nil)
		}
		if u.Message.Text == "/string" {
			api.SendMessage("/token", u.ChatID(), nil)
		}
		if u.Message.Text == "/token" {
			api.SendMessage("token", u.ChatID(), nil)
		}
	}
}
