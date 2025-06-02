package lesson_storage

import "school_schedule_2/internal/domain/model"

func (repo *LessonRepo) CreateLesson(lesson *model.Lesson) (*model.Lesson, error) {
	lesson.ID = repo.nextID
	repo.lessons[repo.nextID] = lesson
	repo.nextID++
	return lesson, nil
}
