package teacher_storage

import (
	"fmt"
	"school_schedule_2/internal/domain/model"
)

func (r *Repo) GetByID(id int64) (*model.Teacher, error) {
	teacher, ok := r.teachers[id]
	if !ok {
		return nil, fmt.Errorf("Teacher not found with identifiler: %d", id)
	}
	return teacher, nil
}
