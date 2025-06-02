package model

import (
	"time"
)

type Teacher struct {
	ID         int64
	Name       string
	Surname    string
	Patronymic string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func NewTeacher(name, surname, patronymic string, now time.Time) *Teacher {
	return &Teacher{
		Name:       name,
		Surname:    surname,
		Patronymic: patronymic,
		CreatedAt:  now,
		UpdatedAt:  now,
	}
}
