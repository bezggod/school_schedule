package service_provider

import (
	"school_schedule_2/internal/controller/class_controller"
	"school_schedule_2/internal/controller/lesson_controller"
	"school_schedule_2/internal/controller/teacher_controller"
)

func (s *ServiceProvider) GetClassController() *class_controller.Controller {
	if s.classController == nil {
		s.classController = class_controller.NewController(s.GetClassUseCase())
	}
	return s.classController
}
func (s *ServiceProvider) GetLessonController() *lesson_controller.Controller {
	if s.lessonController == nil {
		s.lessonController = lesson_controller.NewController(s.GetLessonUseCase())
	}
	return s.lessonController
}

func (s *ServiceProvider) GetTeacherController() *teacher_controller.Controller {
	if s.teacherController == nil {
		s.teacherController = teacher_controller.NewController(s.GetTeacherUseCase())
	}
	return s.teacherController