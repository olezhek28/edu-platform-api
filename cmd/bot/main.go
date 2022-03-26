package main

import (
	"log"

	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/olezhek28/edu-platform-api/internal/app/config"
	"github.com/olezhek28/edu-platform-api/internal/app/service/commander"
	"github.com/olezhek28/edu-platform-api/internal/app/service/lesson"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	bot, err := tgBotAPI.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		log.Fatal(err)
	}

	u := tgBotAPI.UpdateConfig{
		Timeout: 60,
	}

	updates := bot.GetUpdatesChan(u)

	lessonService := lesson.NewLessonService()
	commanderService := commander.NewCommanderService(lessonService, bot, updates)

	err = commanderService.CommandHandler()
	if err != nil {
		log.Fatalf("app crash: %s", err.Error())
	}
}
