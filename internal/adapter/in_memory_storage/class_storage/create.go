package class_storage

import "school_schedule_2/internal/domain/model"

func (r *Repo) CreateClass(class *model.Class) (*model.Class, error) {
	class.ID = r.SetNextID()
	r.classes[class.ID] = class
	return class, nil
}
