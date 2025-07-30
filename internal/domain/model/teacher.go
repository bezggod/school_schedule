package model

import (
	"time"
)

type Teacher struct {
	ID         int64
	Surname    string
	Name       string
	Patronymic string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func NewTeacher(surname, name, patronymic string, now time.Time) *Teacher {
	return &Teacher{
		Surname:    surname,
		Name:       name,
		Patronymic: patronymic,
		CreatedAt:  now,
		UpdatedAt:  now,
	}
}

func (teacher *Teacher) String() string {
	return teacher.Surname + " " + teacher.Name + " " + teacher.Patronymic
}
