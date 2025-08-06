package class_storage

import (
	"school_schedule_2/internal/domain/dto"
	"school_schedule_2/internal/domain/model"
)

func (r *Repo) FindAll(filter dto.FindAllClassesFilter) ([]*model.Class, error) {
	result := make([]*model.Class, 0, len(r.classes))
	for _, class := range r.classes {
		if filter.ClassID != 0 && class.ID != filter.ClassID {
			continue
		}

		if filter.Grade != "" && filter.Grade != class.Grade {
			continue
		}
		result = append(result, class)
	}
	return result, nil
}
