package class_storage

import "school_schedule_2/internal/domain/model"

type Repo struct {
	classes map[int64]*model.Class
	nextID  int64
}

func NewRepo() *Repo {
	return &Repo{
		classes: make(map[int64]*model.Class),
		nextID:  1,
	}
}

func (r *Repo) SetNextID() int64 {
	id := r.nextID
	r.nextID++

	return id
}
