package teacher_storage

import (
	"fmt"
	"school_schedule_2/internal/domain/model"
)

func (r *Repo) UpdateTeacher(teacher *model.Teacher) (*model.Teacher, error) {
	upTeacher, ok := r.teachers[teacher.ID]
	if !ok {
		return nil, fmt.Errorf("teacher with ID %d not found", teacher.ID)
	}

	upTeacher.Name = teacher.Name
	upTeacher.Surname = teacher.Surname
	upTeacher.Patronymic = teacher.Patronymic

	r.teachers[teacher.ID] = upTeacher
	return upTeacher, nil
}
