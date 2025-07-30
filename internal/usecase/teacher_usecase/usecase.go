package teacher_usecase

import (
	"school_schedule_2/internal/domain/dto"
	"school_schedule_2/internal/domain/model"
	"time"
)

type UseCase struct {
	teacherRepo teacherRepo
	timer       timer
}

type (
	teacherRepo interface {
		CreateTeacher(teacher *model.Teacher) (*model.Teacher, error)
		FindAll(req dto.FindAllTeacherFilter) ([]*model.Teacher, error)
		GetByID(id int64) (*model.Teacher, error)
		UpdateTeacher(teacher *model.Teacher) (*model.Teacher, error)
	}

	timer interface {
		NowMoscow() time.Time
	}
)

func NewUseCase(repo teacherRepo, timer timer) *UseCase {
	return &UseCase{
		teacherRepo: repo,
		timer:       timer,
	}
}
