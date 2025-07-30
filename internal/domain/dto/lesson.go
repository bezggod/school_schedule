package dto

import (
	"school_schedule_2/internal/domain/model/enums"
)

type FindAllLessonFilter struct {
	TeacherID  int64
	ClassID    int64
	OfficeName enums.OfficeName
	Subject    enums.SubjectName
	TimeSlot   enums.TimeSlotName
}
