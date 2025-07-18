package lesson_storage

import (
	"fmt"
	"school_schedule_2/internal/domain/model"
)

func (r *Repo) GetByID(id int64) (*model.Lesson, error) {
	lesson, ok := r.lessons[id]
	if !ok {
		return nil, fmt.Errorf("lesson not found")
	}
	return lesson, nil
}
