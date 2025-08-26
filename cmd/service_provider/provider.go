package service_provider

import (
	"school_schedule_2/internal/adapter/postgres/classes"
	"school_schedule_2/internal/adapter/postgres/lessons"
	"school_schedule_2/internal/adapter/postgres/teachers"
	"school_schedule_2/internal/config"
	"school_schedule_2/internal/usecase/class_usecase"
	"school_schedule_2/internal/usecase/lesson_usecase"
	"school_schedule_2/internal/usecase/teacher_usecase"
)

type ServiceProvider struct {
	//repo
	classRepo   *classes.Repo
	teacherRepo *teachers.Repo
	lessonRepo  *lessons.Repo

	// usecase
	classUseCase   *class_usecase.UseCase
	teacherUseCase *teacher_usecase.UseCase
	lessonUseCase  *lesson_usecase.UseCase

	dbCluster *config.Cluster
}

func NewServiceProvider() *ServiceProvider {
	return &ServiceProvider{}
}
