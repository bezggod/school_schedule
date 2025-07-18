package teacher_usecase

import (
	"school_schedule_2/internal/domain/dto"
	"school_schedule_2/internal/domain/model"
)

type UseCase struct {
	TeacherRepo teacherRepo
}

func NewUseCase(repo teacherRepo) *UseCase {
	return &UseCase{
		TeacherRepo: repo,
	}
}

type teacherRepo interface {
	CreateTeacher(teacher *model.Teacher) (*model.Teacher, error)
	FindAll(req dto.FindAllTeacherFilter) ([]*model.Teacher, error)
	GetByID(id int64) (*model.Teacher, error)
	UpdateTeacher(teacher *model.Teacher) (*model.Teacher, error)
}
