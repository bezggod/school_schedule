package teacher_usecase

import (
	"errors"
	"school_schedule_2/internal/domain/model"
	"time"
)

type CreateTeacherReq struct {
	Name       string
	Surname    string
	Patronymic string
}

func (usecase *TeacherUseCase) CreateTeacher(req CreateTeacherReq) (*model.Teacher, error) {
	if req.Name == "" || req.Surname == "" || req.Patronymic == "" {
		return nil, errors.New("Имя, Фамилия, Отчество не могут быть пустыми")
	}

	now := time.Now()
	teacher := model.NewTeacher(req.Name, req.Surname, req.Patronymic, now)

	createdTeacher, err := usecase.TeacherRepo.CreateTeacher(teacher)
	if err != nil {
		return nil, err
	}

	return createdTeacher, nil
}
