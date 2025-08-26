package service_provider

import (
	"context"
	"school_schedule_2/internal/adapter/postgres/lessons"
	"school_schedule_2/internal/adapter/postgres/teachers"

	"school_schedule_2/internal/adapter/postgres/classes"
)

func (s *ServiceProvider) getClassRepo() *classes.Repo {
	if s.classRepo == nil {
		s.classRepo = classes.NewRepo(s.getDbCluster(context.Background()))
	}
	return s.classRepo
}
func (s *ServiceProvider) getTeacherRepo() *teachers.Repo {
	if s.teacherRepo == nil {
		s.teacherRepo = teachers.NewRepo(s.getDbCluster(context.Background()))
	}
	return s.teacherRepo
}

func (s *ServiceProvider) getLessonRepo() *lessons.Repo {
	if s.lessonRepo == nil {
		s.lessonRepo = lessons.NewRepo(s.getDbCluster(context.Background()))
	}
	return s.lessonRepo
}
