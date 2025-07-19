package lesson_storage

import (
	"school_schedule_2/internal/domain/model"
)

type Repo struct {
	lessons map[int64]*model.Lesson
	nextID  int64
}

func NewRepo() *Repo {
	return &Repo{
		lessons: make(map[int64]*model.Lesson),
		nextID:  1,
	}
}

func (r *Repo) SetNextID() int64 {
	id := r.nextID
	r.nextID++

	return id
}
