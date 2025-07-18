package dto

import (
	"school_schedule_2/internal/domain/model/enums"
)

type FindAllTeacherFilter struct {
	TimeSlot enums.TimeSlotName
	Office   enums.OfficeName
}
