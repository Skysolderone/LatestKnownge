package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func init() {
	os.Setenv("TELEGRAM_APITOKEN", "6892958172:AAFCFpZ8RMI4MYpk6E8Vdpiubdly4kyYGtA")
}

func main() {
	bot, err := tgbotapi.NewBotAPI("6892958172:AAFCFpZ8RMI4MYpk6E8Vdpiubdly4kyYGtA") // create new bot
	if err != nil {
		panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.InlineQuery == nil { // if no inline query, ignore it
			continue
		}

		article := tgbotapi.NewInlineQueryResultArticle(update.InlineQuery.ID, "Echo", update.InlineQuery.Query)
		article.Description = update.InlineQuery.Query

		inlineConf := tgbotapi.InlineConfig{
			InlineQueryID: update.InlineQuery.ID,
			IsPersonal:    true,
			CacheTime:     0,
			Results:       []interface{}{article},
		}

		if _, err := bot.Request(inlineConf); err != nil {
			log.Println(err)
		}
	}
}
