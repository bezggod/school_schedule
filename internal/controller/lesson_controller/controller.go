package lesson_controller

import (
	"school_schedule_2/internal/domain/model"
	"school_schedule_2/internal/usecase/lesson_usecase"
)

type (
	Controller struct {
		lessonUseCase lessonUseCase
	}

	lessonUseCase interface {
		Create(req lesson_usecase.CreateLessonReq) (*model.Lesson, error)
		Update(req lesson_usecase.UpdateLessonReq) (*model.Lesson, error)
		FindAll(req lesson_usecase.FindAllLessonReq) ([]*model.Lesson, error)
	}
)

func NewController(lessonUseCase lessonUseCase) *Controller {
	return &Controller{
		lessonUseCase: lessonUseCase,
	}
}
