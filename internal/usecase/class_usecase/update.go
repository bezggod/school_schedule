package class_usecase

import (
	"fmt"

	"school_schedule_2/internal/domain/model"
)

type UpdateClassReq struct {
	ID    int64
	Grade string
}

func (u *UseCase) UpdateClass(req UpdateClassReq) (*model.Class, error) {
	class, err := u.classRepo.GetByID(req.ID)
	if err != nil {
		return nil, fmt.Errorf("classRepo.GetByID: %w", err)
	}

	class.Grade = req.Grade

	class, err = u.classRepo.Update(class)
	if err != nil {
		return nil, fmt.Errorf("classRepo.UpdateClass: %w", err)
	}

	return class, nil
}
