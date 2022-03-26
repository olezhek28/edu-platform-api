package lesson

import (
	"fmt"

	"github.com/olezhek28/edu-platform-api/internal/app/model"
)

func (s *Service) AddLesson(lesson *model.Lesson) (int64, error) {
	lesson.ID = int64(len(s.lessons)) + 1
	s.lessons[lesson.ID] = lesson

	fmt.Println("lesson added")
	return 0, nil
}
