package teacher_storage

import "school_schedule_2/internal/domain/model"

func (r *TeacherRepo) CreateTeacher(teacher *model.Teacher) (*model.Teacher, error) {
	teacher.ID = r.SetNextID()
	r.teachers[r.nextID] = teacher
	return teacher, nil
}
