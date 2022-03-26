package lesson

import (
	"github.com/olezhek28/edu-platform-api/internal/app/model"
)

type Service struct {
	lessons map[int64]*model.Lesson
}

func NewLessonService() *Service {
	return &Service{
		lessons: make(map[int64]*model.Lesson),
	}
}
