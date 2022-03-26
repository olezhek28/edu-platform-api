package commander

import (
	"fmt"
	"strconv"

	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/olezhek28/edu-platform-api/internal/app/model"
	"github.com/olezhek28/edu-platform-api/internal/app/service/lesson"
)

type Service struct {
	lessonService *lesson.Service

	bot     *tgBotAPI.BotAPI
	updates tgBotAPI.UpdatesChannel

	commands map[string]func()
}

func NewCommanderService(
	lessonService *lesson.Service,

	bot *tgBotAPI.BotAPI,
	updates tgBotAPI.UpdatesChannel,
) *Service {
	return &Service{
		lessonService: lessonService,

		bot:     bot,
		updates: updates,
	}
}

func (s *Service) CommandHandler() error {
	for update := range s.updates {
		if update.Message != nil { // If we got a message
			switch update.Message.Command() {
			case "add":
				id, err := s.lessonService.AddLesson(&model.Lesson{
					ID:          0,
					Title:       "Test",
					Description: "blabla",
					Duration:    12,
				})
				if err != nil {
					return nil
				}
				msg := tgBotAPI.NewMessage(update.Message.Chat.ID, strconv.FormatInt(id, 10))
				s.bot.Send(msg)
			case "delete":
				err := s.lessonService.DeleteLesson(1)
				if err != nil {
					return nil
				}
			case "list":
				list, err := s.lessonService.GetListLesson()
				if err != nil {
					return nil
				}
				var listStr string
				for _, l := range list {
					listStr += fmt.Sprintf("%#v\n", l)
				}

				msg := tgBotAPI.NewMessage(update.Message.Chat.ID, listStr)
				s.bot.Send(msg)
			case "update":
				err := s.lessonService.UpdateLesson(&model.Lesson{
					ID:          1,
					Title:       "Truck",
					Description: "kluklu",
					Duration:    21,
				})
				if err != nil {
					return nil
				}
			}
		}
	}

	return nil
}
