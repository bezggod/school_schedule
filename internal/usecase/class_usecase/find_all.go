package class_usecase

import (
	"fmt"
	"school_schedule_2/internal/domain/dto"
	"school_schedule_2/internal/domain/model"
)

type FindAllReq struct {
	ClassID int64
	Grade   string
}

func (u *UseCase) FindAll(req FindAllReq) ([]*model.Class, error) {
	classes, err := u.classRepo.FindAll(dto.FindAllClassesFilter{
		ClassID: req.ClassID,
		Grade:   req.Grade,
	})
	if err != nil {
		return nil, fmt.Errorf("classRepo.FindAll: %w", err)
	}
	return classes, nil
}
