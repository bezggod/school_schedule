package teacher_controller

import (
	"school_schedule_2/internal/domain/model"
	"school_schedule_2/internal/usecase/teacher_usecase"
)

type (
	Controller struct {
		teacherUseCase teacherUseCase
	}

	teacherUseCase interface {
		Create(req teacher_usecase.CreateTeacherReq) (*model.Teacher, error)
		Update(req teacher_usecase.UpdateTeacherReq) (*model.Teacher, error)
	}
)

func NewController(teacherUseCase teacherUseCase) *Controller {
	return &Controller{
		teacherUseCase: teacherUseCase,
	}
}
