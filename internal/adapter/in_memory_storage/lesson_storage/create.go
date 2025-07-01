package lesson_storage

import "school_schedule_2/internal/domain/model"

func (r *LessonRepo) CreateLesson(lesson *model.Lesson) (*model.Lesson, error) {
	lesson.ID = r.SetNextID()
	r.lessons[lesson.ID] = lesson
	return lesson, nil
}
