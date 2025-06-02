package teacher_storage

import "school_schedule_2/internal/domain/model"

func (repo *TeacherRepo) CreateTeacher(teacher *model.Teacher) (*model.Teacher, error) {
	teacher.ID = repo.nextID
	repo.teachers[repo.nextID] = teacher
	repo.nextID++
	return teacher, nil
}
