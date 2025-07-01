package class_usecase

import (
	"school_schedule_2/internal/adapter/in_memory_storage/class_storage"
)

type UseCase struct {
	ClassRepo *class_storage.ClassRepo
}

func NewUseCase(repo *class_storage.ClassRepo) *UseCase {
	return &UseCase{
		ClassRepo: repo,
	}
}
