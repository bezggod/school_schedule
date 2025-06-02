package model

import (
	"school_schedule_2/internal/domain/model/enums"
	"time"
)

type Office struct {
	ID        int64
	Name      enums.OfficeName
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewOffice(name enums.OfficeName, now time.Time) *Office {
	return &Office{
		Name:      name,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
