package lesson_storage

import "school_schedule_2/internal/domain/model"

type LessonRepo struct {
	lessons map[int64]*model.Lesson
	nextID  int64
}

func NewLessonRepo() *LessonRepo {
	return &LessonRepo{
		lessons: make(map[int64]*model.Lesson),
		nextID:  1,
	}
}
