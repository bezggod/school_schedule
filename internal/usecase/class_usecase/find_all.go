package class_usecase

import (
	"fmt"
	"school_schedule_2/internal/domain/dto"
	"school_schedule_2/internal/domain/model"
	"school_schedule_2/internal/domain/model/enums"
)

type FindAllReq struct {
	ClassID  int64
	TimeSlot enums.TimeSlotName
	Office   enums.OfficeName
}

func (u *UseCase) FindAll(req FindAllReq) ([]*model.Class, error) {
	classes, err := u.ClassRepo.FindAll(dto.FindAllClassesFilter{
		ClassID:  req.ClassID,
		TimeSlot: req.TimeSlot,
		Office:   req.Office,
	})
	if err != nil {
		return nil, fmt.Errorf("ClassRepo.FindAll: %w", err)
	}
	return classes, nil
}
