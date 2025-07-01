package lesson_storage

import "school_schedule_2/internal/domain/model"

func (r *LessonRepo) GetAll() ([]*model.Lesson, error) {
	result := make([]*model.Lesson, 0, len(r.lessons))
	for _, lesson := range r.lessons {
		result = append(result, lesson)
	}
	return result, nil
}
