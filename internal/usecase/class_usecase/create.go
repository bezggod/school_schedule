package class_usecase //тут создание новых классов

import (
	"errors"
	"fmt"
	"school_schedule_2/internal/domain/model"
	"time"
)

type CreateClassReq struct {
	Grade string
}

func (usecase *ClassUseCase) CreateClass(req CreateClassReq) (*model.Class, error) {
	if req.Grade == "" {
		return nil, errors.New("Grade пуст")
	}

	now := time.Now()
	class := model.NewClass(req.Grade, now)

	class, err := usecase.ClassRepo.CreateClass(class)
	if err != nil {
		return nil, fmt.Errorf("classRepo.CreateClass: %w", err)
	}
	return class, nil
}
