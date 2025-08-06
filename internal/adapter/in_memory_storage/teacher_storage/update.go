package teacher_storage

import (
	"school_schedule_2/internal/domain/model"
)

func (r *Repo) Update(teacher *model.Teacher) (*model.Teacher, error) {
	r.teachers[teacher.ID] = teacher

	return teacher, nil
}
