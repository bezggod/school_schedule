package class_controller

import (
	"school_schedule_2/internal/domain/model"
	"school_schedule_2/internal/usecase/class_usecase"
)

type (
	Controller struct {
		classUseCase classUseCase
	}

	classUseCase interface {
		Create(req class_usecase.CreateClassReq) (*model.Class, error)
		Update(req class_usecase.UpdateClassReq) (*model.Class, error)
		FindAll(req class_usecase.FindAllClassReq) ([]*model.Class, error)
	}
)

func NewController(classUseCase classUseCase) *Controller {
	return &Controller{
		classUseCase: classUseCase,
	}
}
