package teacher_storage

import (
	"fmt"
	"school_schedule_2/internal/domain/model"
)

func (r *TeacherRepo) GetByID(id int64) (*model.Teacher, error) {
	teacher, ok := r.teachers[id]
	if !ok {
		return nil, fmt.Errorf("Учитель не найден", id)
	}
	return teacher, nil
}
