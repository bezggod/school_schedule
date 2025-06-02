package lesson_usecase

import (
	"school_schedule_2/internal/adapter/in_memory_storage/lesson_storage"
)

type LessonUseCase struct {
	LessonRepo *lesson_storage.LessonRepo
}

func NewLessonUseCase(repo *lesson_storage.LessonRepo) *LessonUseCase {
	return &LessonUseCase{
		LessonRepo: repo,
	}
}
