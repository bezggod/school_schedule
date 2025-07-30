package teacher_storage

import (
	"fmt"
	"school_schedule_2/internal/domain/dto"
	"school_schedule_2/internal/domain/model"
)

func (r *Repo) UpdateTeacher(filter dto.UpdateTeacherFilter) (*model.Teacher, error) {
	upTeacher, ok := r.teachers[filter.ID]
	if !ok {
		return nil, fmt.Errorf("teacher with ID %d not found", filter.ID)
	}

	if filter.Name != nil {
		upTeacher.Name = *filter.Name
	}

	if filter.Surname != nil {
		upTeacher.Surname = *filter.Surname
	}
	if filter.Patronymic != nil {
		upTeacher.Patronymic = *filter.Patronymic
	}

	r.teachers[filter.ID] = upTeacher
	return upTeacher, nil
}
