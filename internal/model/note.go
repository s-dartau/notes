package model

import "time"

type Note struct {
	ID        uint64
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
