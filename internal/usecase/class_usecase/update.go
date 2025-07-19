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
	class, err := u.ClassRepo.GetByID(req.ID)
	if err != nil {
		return nil, fmt.Errorf("get class by id err: %w", err)
	}

	class.Grade = req.Grade

	update, err := u.ClassRepo.UpdateClass(class)
	if err != nil {
		return nil, fmt.Errorf("update class: %w", err)
	}
	return update, nil
}
