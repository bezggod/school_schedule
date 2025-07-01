package class_storage

import "school_schedule_2/internal/domain/model"

type ClassRepo struct {
	classes map[int64]*model.Class
	nextID  int64
}

func NewClassRepo() *ClassRepo {
	return &ClassRepo{
		classes: make(map[int64]*model.Class),
		nextID:  1,
	}
}

func (r *ClassRepo) SetNextID() int64 {
	id := r.nextID
	r.nextID++

	return id
}
