package class_usecase //тут создание новых классов

import (
	"fmt"
	"school_schedule_2/internal/domain/model"
	"time"
)

type CreateClassReq struct {
	Grade string
}

func (u *UseCase) CreateClass(req CreateClassReq) (*model.Class, error) {
	now := time.Now()
	class := model.NewClass(req.Grade, now)

	createClass, err := u.ClassRepo.CreateClass(class)
	if err != nil {
		return nil, fmt.Errorf("classRepo.CreateClass: %w", err)
	}
	return createClass, nil
}
