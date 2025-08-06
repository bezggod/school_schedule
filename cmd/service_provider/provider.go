package service_provider

import (
	"school_schedule_2/internal/adapter/in_memory_storage/class_storage"
	"school_schedule_2/internal/adapter/in_memory_storage/lesson_storage"
	"school_schedule_2/internal/adapter/in_memory_storage/teacher_storage"
	"school_schedule_2/internal/config"
	"school_schedule_2/internal/usecase/class_usecase"
	"school_schedule_2/internal/usecase/lesson_usecase"
	"school_schedule_2/internal/usecase/teacher_usecase"
)

type ServiceProvider struct {
	//repo
	classRepo   *class_storage.Repo
	teacherRepo *teacher_storage.Repo
	lessonRepo  *lesson_storage.Repo

	// usecase
	classUseCase   *class_usecase.UseCase
	teacherUseCase *teacher_usecase.UseCase
	lessonUseCase  *lesson_usecase.UseCase

	dbCluster *config.Cluster
}

func NewServiceProvider() *ServiceProvider {
	return &ServiceProvider{}
}
