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

func (ut *UseCase) UpdateTeacher(req UpdateTeacherReq) (*model.Teacher, error) {
	teacher, err := ut.TeacherRepo.GetByID(req.ID)
	if err != nil {
		return nil, fmt.Errorf("find teacher by id: %w", err)
	}

	teacher.Name = req.Name
	teacher.Surname = req.Surname
	teacher.Patronymic = req.Patronymic

	update, err := ut.TeacherRepo.UpdateTeacher(teacher)
	if err != nil {
		return nil, fmt.Errorf("update teacher: %w", err)
	}
	return update, nil
}
