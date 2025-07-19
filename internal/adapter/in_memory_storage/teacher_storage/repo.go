package teacher_storage

import "school_schedule_2/internal/domain/model"

type Repo struct {
	teachers map[int64]*model.Teacher
	nextID   int64
}

func NewTeacherRepo() *Repo {
	return &Repo{
		teachers: make(map[int64]*model.Teacher),
		nextID:   1,
	}
}
func (r *Repo) SetNextID() int64 {
	id := r.nextID
	r.nextID++

	return id
}
