package class_storage

import (
	"fmt"

	"school_schedule_2/internal/domain/model"
)

func (r *Repo) GetByID(id int64) (*model.Class, error) {
	class, ok := r.classes[id]
	if !ok {
		return nil, fmt.Errorf("classes not found")
	}

	return class, nil
}
