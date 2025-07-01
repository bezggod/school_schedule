package teacher_usecase

import (
	"school_schedule_2/internal/domain/model"
	"time"
)

type CreateTeacherReq struct {
	Name       string
	Surname    string
	Patronymic string
}

func (u *TeacherUseCase) CreateTeacher(req CreateTeacherReq) (*model.Teacher, error) {

	now := time.Now()
	teacher := model.NewTeacher(req.Name, req.Surname, req.Patronymic, now)

	createdTeacher, err := u.TeacherRepo.CreateTeacher(teacher)
	if err != nil {
		return nil, err
	}

	return createdTeacher, nil
}
