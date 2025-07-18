package teacher_usecase

import (
	"fmt"
	"school_schedule_2/internal/domain/model"
	"time"
)

type CreateTeacherReq struct {
	Name       string
	Surname    string
	Patronymic string
}

func (ut *UseCase) CreateTeacher(req CreateTeacherReq) (*model.Teacher, error) {

	now := time.Now()
	teacher := model.NewTeacher(req.Name, req.Surname, req.Patronymic, now)

	createdTeacher, err := ut.TeacherRepo.CreateTeacher(teacher)
	if err != nil {
		return nil, fmt.Errorf("ut.TeacherRepo.CreateTeacher: %w", err)
	}

	return createdTeacher, nil
}
