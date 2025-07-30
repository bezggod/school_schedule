package teacher_usecase

import (
	"fmt"
	"school_schedule_2/internal/domain/dto"
	"school_schedule_2/internal/domain/model"
)

type FindAllReq struct {
	Surname    string
	Name       string
	Patronymic string
}

func (ut *UseCase) FindAll(req FindAllReq) ([]*model.Teacher, error) {
	teachers, err := ut.teacherRepo.FindAll(dto.FindAllTeacherFilter{
		Surname: req.Surname,
		Name:    req.Name,
	})
	if err != nil {
		return nil, fmt.Errorf("teacherRepo.FindAll: %w", err)
	}
	return teachers, nil

}
