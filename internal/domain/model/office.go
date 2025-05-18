package model

import "time"

type Office struct {
	ID        int64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewOffice(name string, now time.Time) *Office {
	return &Office{
		Name:      name,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
