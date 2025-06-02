package teacher_usecase

import (
	"school_schedule_2/internal/adapter/in_memory_storage/teacher_storage"
)

type TeacherUseCase struct {
	TeacherRepo *teacher_storage.TeacherRepo
}

func NewTeacherUseCase(repo *teacher_storage.TeacherRepo) *TeacherUseCase {
	return &TeacherUseCase{
		TeacherRepo: repo,
	}
}
