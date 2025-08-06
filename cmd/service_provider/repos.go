package service_provider

import (
	"school_schedule_2/internal/adapter/in_memory_storage/class_storage"
	"school_schedule_2/internal/adapter/in_memory_storage/lesson_storage"
	"school_schedule_2/internal/adapter/in_memory_storage/teacher_storage"
)

func (s *ServiceProvider) getClassRepo() *class_storage.Repo {
	if s.classRepo == nil {
		s.classRepo = class_storage.NewRepo()
	}
	return s.classRepo
}
func (s *ServiceProvider) getTeacherRepo() *teacher_storage.Repo {
	if s.teacherRepo == nil {
		s.teacherRepo = teacher_storage.NewRepo()
	}
	return s.teacherRepo
}

func (s *ServiceProvider) getLessonRepo() *lesson_storage.Repo {
	if s.lessonRepo == nil {
		s.lessonRepo = lesson_storage.NewRepo()
	}
	return s.lessonRepo
}
