package model

import (
	"time"
)

type Office struct {
	ID        int64
	Name      uint8
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewOffice(name uint8, now time.Time) *Office {
	return &Office{
		Name:      name,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
