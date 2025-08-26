package class_usecase

import (
	"fmt"
	"school_schedule_2/internal/domain/model"
)

type CreateClassReq struct {
	Grade string
}

func (u *UseCase) CreateClass(req CreateClassReq) (*model.Class, error) {

	now := u.timer.NowMoscow()
	class := model.NewClass(req.Grade, now)

	createdClass, err := u.classRepo.Create(class)
	if err != nil {
		return nil, fmt.Errorf("classRepo.Create: %w", err)
	}

	return createdClass, nil
}
