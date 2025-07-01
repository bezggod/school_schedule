package lesson_storage

import (
	"school_schedule_2/internal/domain/model"
	"school_schedule_2/internal/domain/model/enums"
)

type LessonRepo struct {
	lessons map[int64]*model.Lesson
	nextID  int64
}

func NewClassRepo() *LessonRepo {
	return &LessonRepo{
		lessons: make(map[int64]*model.Lesson),
		nextID:  1,
	}
}

func (r *LessonRepo) SetNextID() int64 {
	id := r.nextID
	r.nextID++

	return id
}

func (r *LessonRepo) LessonExists(name enums.OfficeName, slot model.TimeSlot) bool {
	for _, lesson := range r.lessons {
		if lesson.Office.Name == name && lesson.TimeSlot.Slot == slot.Slot {
			return true
		}
	}
	return false
}
