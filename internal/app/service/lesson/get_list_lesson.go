package lesson

import (
	"fmt"

	"github.com/olezhek28/edu-platform-api/internal/app/model"
)

func (s *Service) GetListLesson() ([]*model.Lesson, error) {
	res := make([]*model.Lesson, 0, len(s.lessons))
	for _, l := range s.lessons {
		fmt.Println(l)
		res = append(res, l)
	}

	fmt.Println("lesson list...")
	return res, nil
}
