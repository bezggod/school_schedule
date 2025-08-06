package service_provider

import (
	"school_schedule_2/internal/pkg"
	"school_schedule_2/internal/usecase/class_usecase"
	"school_schedule_2/internal/usecase/lesson_usecase"
	"school_schedule_2/internal/usecase/teacher_usecase"
)

func (s *ServiceProvider) GetLessonUseCase() *lesson_usecase.UseCase {
	if s.lessonUseCase == nil {
		s.lessonUseCase = lesson_usecase.NewUseCase(s.getLessonRepo(), s.getTeacherRepo(), s.getClassRepo(), pkg.NewTimer())
	}
	return s.lessonUseCase
}
func (s *ServiceProvider) GetClassUseCase() *class_usecase.UseCase {
	if s.classUseCase == nil {
		s.classUseCase = class_usecase.NewUseCase(s.getClassRepo(), pkg.NewTimer())
	}
	return s.classUseCase
}
func (s *ServiceProvider) GetTeacherUseCase() *teacher_usecase.UseCase {
	if s.teacherUseCase == nil {
		s.teacherUseCase = teacher_usecase.NewUseCase(s.getTeacherRepo(), pkg.NewTimer())
	}
	return s.teacherUseCase
}
