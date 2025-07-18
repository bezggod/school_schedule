package lesson_storage

import (
	"school_schedule_2/internal/domain/dto"
	"school_schedule_2/internal/domain/model"
)

func (r *Repo) FindLesson(filter dto.FindAllLessonFilter) ([]*model.Lesson, error) {
	result := make([]*model.Lesson, 0, len(r.lessons))
	for _, lesson := range r.lessons {
		if filter.TeacherID != 0 && filter.TeacherID != lesson.TeacherID { //пропуск учителя, если он у
			continue
		}

		if filter.OfficeName != "" && filter.OfficeName != lesson.Office {
			continue
		}
		result = append(result, lesson)
	}

	return result, nil
}
