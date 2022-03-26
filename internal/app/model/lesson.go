package model

import "time"

type Lesson struct {
	ID          int64
	Title       string
	Description string
	Duration    time.Duration
}
