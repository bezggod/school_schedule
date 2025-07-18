package teacher_storage

import "school_schedule_2/internal/domain/model"

func (r *Repo) CreateTeacher(teacher *model.Teacher) (*model.Teacher, error) {
	teacher.ID = r.SetNextID()
	r.teachers[teacher.ID] = teacher
	return teacher, nil
}
