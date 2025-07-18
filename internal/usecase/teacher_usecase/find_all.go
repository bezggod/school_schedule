package teacher_usecase

import (
	"fmt"
	"school_schedule_2/internal/domain/dto"
	"school_schedule_2/internal/domain/model"
	"school_schedule_2/internal/domain/model/enums"
)

type FindAllReq struct {
	TimeSlot enums.TimeSlotName
	Office   enums.OfficeName
}

func (ut *UseCase) FindAll(req FindAllReq) ([]*model.Teacher, error) {
	teachers, err := ut.TeacherRepo.FindAll(dto.FindAllTeacherFilter{
		TimeSlot: req.TimeSlot,
		Office:   req.Office,
	})
	if err != nil {
		return nil, fmt.Errorf("TeacherRepo.FindAll: %w", err)
	}
	return teachers, nil

}
