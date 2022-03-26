package lesson

import (
	"fmt"

	"github.com/olezhek28/edu-platform-api/internal/app/model"
)

func (s *Service) UpdateLesson(lesson *model.Lesson) error {
	if _, found := s.lessons[lesson.ID]; !found {
		return nil
	}
	s.lessons[lesson.ID] = lesson

	fmt.Println("lesson updated")
	return nil
}
