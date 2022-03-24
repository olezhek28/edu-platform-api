package main

import (
	"log"

	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/olezhek28/edu-platform-api/internal/app/config"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	bot, err := tgBotAPI.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgBotAPI.UpdateConfig{
		Timeout: 60,
	}

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgBotAPI.NewMessage(update.Message.Chat.ID, "You wrote:"+update.Message.Text)

			bot.Send(msg)
		}
	}
}
