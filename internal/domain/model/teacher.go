package model

import "time"

type Teacher struct {
	ID        int64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewTeacher(name string, now time.Time) *Teacher {
	return &Teacher{
		Name:      name,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
