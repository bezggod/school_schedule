package lesson_storage

import (
	"school_schedule_2/internal/domain/model/enums"
)

func (r *Repo) LessonExists(officeName enums.OfficeName, slot enums.TimeSlotName) bool {
	for _, lesson := range r.lessons {
		if lesson.Office == officeName && lesson.TimeSlot == slot {
			return true
		}
	}
	return false
}
