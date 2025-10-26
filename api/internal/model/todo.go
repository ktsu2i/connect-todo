package model

import "time"

type Todo struct {
	ID        int64
	Title     string
	Done      bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
