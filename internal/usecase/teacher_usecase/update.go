package teacher_usecase

import (
	"fmt"
	"school_schedule_2/internal/domain/model"
)

type UpdateTeacherReq struct {
	ID         int64
	Name       string
	Patronymic string
	Surname    string
}

func (ut *UseCase) Update(req UpdateTeacherReq) (*model.Teacher, error) {
	teacher, err := ut.teacherRepo.GetByID(req.ID)
	if err != nil {
		return nil, fmt.Errorf("teacherRepo.GetByID: %w", err)
	}

	teacher.Surname = req.Surname
	teacher.Name = req.Name
	teacher.Patronymic = req.Patronymic

	teacher, err = ut.teacherRepo.Update(teacher)
	if err != nil {
		return nil, fmt.Errorf("teacherRepo.Update: %w", err)
	}

	return teacher, nil
}
