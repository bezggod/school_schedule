package class_usecase

import (
	"time"

	"school_schedule_2/internal/domain/dto"
	"school_schedule_2/internal/domain/model"
)

type UseCase struct {
	classRepo classRepo
	timer     timer
}

type (
	classRepo interface {
		CreateClass(class *model.Class) (*model.Class, error)
		FindAll(req dto.FindAllClassesFilter) ([]*model.Class, error)
		GetByID(id int64) (*model.Class, error)
		Update(class *model.Class) (*model.Class, error)
	}

	timer interface {
		NowMoscow() time.Time
	}
)

func NewUseCase(repo classRepo, timer timer) *UseCase {
	return &UseCase{
		classRepo: repo,
		timer:     timer,
	}
}
