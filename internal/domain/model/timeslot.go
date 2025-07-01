package model

import (
	"school_schedule_2/internal/domain/model/enums"
	"time"
)

type TimeSlot struct {
	ID        int64
	Slot      enums.TimeSlotName
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewTimeSlot(name enums.TimeSlotName, now time.Time) *TimeSlot {
	return &TimeSlot{
		Slot:      name,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
