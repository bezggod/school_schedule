package office_storage

import "school_schedule_2/internal/domain/model"

type OfficeRepo struct {
	offices map[int64]*model.Office
	nextID  int64
}

func NewOfficeRepo() *OfficeRepo {
	return &OfficeRepo{
		offices: make(map[int64]*model.Office),
		nextID:  1,
	}
}
