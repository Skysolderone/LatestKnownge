package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	bot     *tgbotapi.BotAPI
	updates tgbotapi.UpdatesChannel
)

func main() {
	var err error
	bot, err = tgbotapi.NewBotAPI("6892958172:AAFCFpZ8RMI4MYpk6E8Vdpiubdly4kyYGtA")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates = bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			handleMessage(update.Message)
		}
	}
}

func handleMessage(msg *tgbotapi.Message) {
	switch msg.Text {
	case "/start":
		sendMessage(msg.Chat.ID, "Welcome! Please enter your name:")
		waitForUserInput(msg.Chat.ID, handleNameInput)
	default:
		sendMessage(msg.Chat.ID, "I don't understand that command.")
	}
}

func sendMessage(chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	bot.Send(msg)
}

func waitForUserInput(chatID int64, callback func(*tgbotapi.Message)) {
	for update := range updates {
		if update.Message != nil && update.Message.Chat.ID == chatID {
			callback(update.Message)
			break
		}
	}
}

func handleNameInput(msg *tgbotapi.Message) {
	sendMessage(msg.Chat.ID, "Hello, "+msg.Text+"! How can I assist you today?")
}
