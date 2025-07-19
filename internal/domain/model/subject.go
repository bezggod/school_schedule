package model

import (
	"school_schedule_2/internal/domain/model/enums"
	"time"
)

type Subject struct {
	ID        int64
	Name      enums.SubjectName
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewSubject(name enums.SubjectName, now time.Time) *Subject {
	return &Subject{
		Name:      name,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
