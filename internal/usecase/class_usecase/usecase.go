package class_usecase

import (
	"school_schedule_2/internal/domain/dto"
	"school_schedule_2/internal/domain/model"
)

type UseCase struct {
	ClassRepo classRepo
}

func NewUseCase(repo classRepo) *UseCase {
	return &UseCase{
		ClassRepo: repo,
	}
}

type classRepo interface {
	CreateClass(class *model.Class) (*model.Class, error)
	FindAll(req dto.FindAllClassesFilter) ([]*model.Class, error)
	GetByID(id int64) (*model.Class, error)
	UpdateClass(class *model.Class) (*model.Class, error)
}
