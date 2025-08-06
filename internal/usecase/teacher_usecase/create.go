package teacher_usecase

import (
	"fmt"
	"school_schedule_2/internal/domain/model"
)

type CreateTeacherReq struct {
	Surname    string
	Name       string
	Patronymic string
}

func (ut *UseCase) CreateTeacher(req CreateTeacherReq) (*model.Teacher, error) {

	now := ut.timer.NowMoscow()
	teacher := model.NewTeacher(req.Surname, req.Name, req.Patronymic, now)

	createdTeacher, err := ut.teacherRepo.CreateTeacher(teacher)
	if err != nil {
		return nil, fmt.Errorf("ut.teacherRepo.CreateTeacher: %w", err)
	}

	return createdTeacher, nil
}
