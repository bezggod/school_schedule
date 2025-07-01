package teacher_storage

import "school_schedule_2/internal/domain/model"

func (r *TeacherRepo) GetAll() ([]*model.Teacher, error) {
	result := make([]*model.Teacher, 0, len(r.teachers))
	for _, teacher := range r.teachers {
		result = append(result, teacher)
	}
	return result, nil
}
