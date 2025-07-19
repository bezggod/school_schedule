package dto

import "school_schedule_2/internal/domain/model/enums"

type FindAllClassesFilter struct {
	ClassID  int64
	TimeSlot enums.TimeSlotName
	Office   enums.OfficeName
}
