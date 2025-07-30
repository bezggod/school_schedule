package lesson_usecase

import (
	"school_schedule_2/internal/domain/dto"
	"school_schedule_2/internal/domain/model"
	"school_schedule_2/internal/domain/model/enums"
	"time"
)

type UseCase struct {
	lessonRepo  lessonRepo
	teacherRepo teacherRepo
	classRepo   classRepo
	timer       timer
}

type (
	lessonRepo interface {
		CreateLesson(lesson *model.Lesson) (*model.Lesson, error)
		LessonExists(name enums.OfficeName, slot enums.TimeSlotName) bool
		FindLesson(req dto.FindAllLessonFilter) ([]*model.Lesson, error)
		GetByID(id int64) (*model.Lesson, error)
	}

	teacherRepo interface {
		GetByID(id int64) (*model.Teacher, error)
	}

	classRepo interface {
		GetByID(id int64) (*model.Class, error)
	}

	timer interface {
		NowMoscow() time.Time
	}
)

func NewUseCase(lessonRepo lessonRepo, teacherRepo teacherRepo, classRepo classRepo, timer timer) *UseCase {
	return &UseCase{
		lessonRepo:  lessonRepo,
		teacherRepo: teacherRepo,
		classRepo:   classRepo,
		timer:       timer,
	}
}
