package model

import (
	"time"
)

type Class struct {
	ID        int64
	Grade     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewClass(grade string, now time.Time) *Class {
	return &Class{
		Grade:     grade,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
