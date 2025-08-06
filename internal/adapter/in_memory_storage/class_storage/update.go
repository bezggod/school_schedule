package class_storage

import (
	"school_schedule_2/internal/domain/model"
)

func (r *Repo) Update(class *model.Class) (*model.Class, error) {
	r.classes[class.ID] = class

	return class, nil
}
