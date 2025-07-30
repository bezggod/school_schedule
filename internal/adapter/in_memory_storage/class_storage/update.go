package class_storage

import (
	"fmt"
	"school_schedule_2/internal/domain/dto"
	"school_schedule_2/internal/domain/model"
	"time"
)

func (r *Repo) UpdateClass(filter dto.UpdateClassFilter) (*model.Class, error) {
	upClass, ok := r.classes[filter.ClassID]
	if !ok {
		return nil, fmt.Errorf("class with ID %d not found", filter.ClassID)
	}
	if filter.Grade != nil && upClass.Grade != *filter.Grade {
		upClass.Grade = *filter.Grade
	}
	upClass.UpdatedAt = time.Now()
	return upClass, nil
}
