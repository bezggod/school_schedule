package lesson_usecase

import (
	"school_schedule_2/internal/domain/model"
	"school_schedule_2/internal/domain/model/enums"
)

type LessonUseCase struct {
	lessonRepo  lessonRepo
	teacherRepo teacherRepo
	classRepo   classRepo
}

func NewLessonUseCase(lessonRepo lessonRepo, teacherRepo teacherRepo, classRepo classRepo) *LessonUseCase {
	return &LessonUseCase{
		lessonRepo:  lessonRepo,
		teacherRepo: teacherRepo,
		classRepo:   classRepo,
	}
}

type lessonRepo interface {
	CreateLesson(lesson *model.Lesson) (*model.Lesson, error)
	LessonExists(name enums.OfficeName, slot model.TimeSlot) bool
}

type teacherRepo interface {
	GetByID(id int64) (*model.Teacher, error)
}

type classRepo interface {
	GetByID(id int64) (*model.Class, error)
}
