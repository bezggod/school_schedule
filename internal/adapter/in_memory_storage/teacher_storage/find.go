package teacher_storage

import (
	"school_schedule_2/internal/domain/dto"
	"school_schedule_2/internal/domain/model"
)

func (r *Repo) FindAll(filter dto.FindAllTeacherFilter) ([]*model.Teacher, error) {
	result := make([]*model.Teacher, 0, len(r.teachers))
	for _, teacher := range r.teachers {
		if filter.Surname != "" && teacher.Surname != filter.Surname {
			continue
		}
		if filter.Name != "" && teacher.Name != filter.Name {
			continue
		}

		result = append(result, teacher)
	}
	return result, nil
}
