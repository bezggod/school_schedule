package class_usecase

import (
	"school_schedule_2/internal/adapter/in_memory_storage/class_storage"
)

type ClassUseCase struct {
	ClassRepo *class_storage.ClassRepo
}

func NewClassUseCase(repo *class_storage.ClassRepo) *ClassUseCase {
	return &ClassUseCase{
		ClassRepo: repo,
	}
}
