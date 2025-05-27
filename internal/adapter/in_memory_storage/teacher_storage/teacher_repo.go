package teacher_storage

import "school_schedule_2/internal/domain/model"

type TeacherRepo struct {
	teachers map[int64]*model.Teacher
	nextID   int64
}

func NewTeacherRepo() *TeacherRepo {
	return &TeacherRepo{
		teachers: make(map[int64]*model.Teacher),
		nextID:   1,
	}
}
