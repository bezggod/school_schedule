package teacher_usecase

import (
	"fmt"
	"school_schedule_2/internal/domain/model"
)

type UpdateTeacherReq struct {
	ID         int64
	Name       *string
	Patronymic *string
	Surname    *string
}

func (ut *UseCase) UpdateTeacher(req UpdateTeacherReq) (*model.Teacher, error) {
	if req.ID == 0 {
		return nil, fmt.Errorf("invalid teacher id")
	}

	teacher, err := ut.teacherRepo.GetByID(req.ID)
	if err != nil {
		return nil, fmt.Errorf("teacherRepo.GetByID: %w", err)
	}

	if req.Name != nil {
		teacher.Name = *req.Name
	}
	if req.Surname != nil {
		teacher.Surname = *req.Surname
	}
	if req.Patronymic != nil {
		teacher.Patronymic = *req.Patronymic
	}

	update, err := ut.teacherRepo.UpdateTeacher(teacher)
	if err != nil {
		return nil, fmt.Errorf("teacherRepo.UpdateTeacher: %w", err)
	}

	return update, nil
}
