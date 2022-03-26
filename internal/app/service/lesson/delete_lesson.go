package lesson

import "fmt"

func (s *Service) DeleteLesson(id int64) error {
	delete(s.lessons, id)

	fmt.Println("lesson deleted")
	return nil
}
