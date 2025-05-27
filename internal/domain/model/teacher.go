package model

import (
	"school_schedule_2/internal/domain/model/enums/teacher_name"
	"time"
)

type Teacher struct {
	ID        int64
	Name      teacher_name.TeacherName
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewTeacher(name teacher_name.TeacherName, now time.Time) *Teacher {
	return &Teacher{
		Name:      name,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
