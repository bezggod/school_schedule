package class_usecase

import (
	"fmt"
	"school_schedule_2/internal/domain/dto"
	"school_schedule_2/internal/domain/model"
)

type UpdateClassReq struct {
	ID    int64
	Grade *string
}

func (u *UseCase) UpdateClass(req UpdateClassReq) (*model.Class, error) {
	if req.ID == 0 {
		return nil, fmt.Errorf("invalid class ID")
	}

	filter := dto.UpdateClassFilter{
		ClassID: req.ID,
		Grade:   req.Grade,
	}

	class, err := u.classRepo.UpdateClass(filter)
	if err != nil {
		return nil, fmt.Errorf("classRepo.UpdateClass: %w", err)
	}
	return class, nil
}
