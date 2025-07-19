package class_storage

import (
	"fmt"
	"school_schedule_2/internal/domain/model"
)

func (r *Repo) UpdateClass(class *model.Class) (*model.Class, error) {
	upClass, ok := r.classes[class.ID]
	if !ok {
		return nil, fmt.Errorf("class with ID %d not found", class.ID)
	}

	upClass.Grade = class.Grade

	r.classes[class.ID] = upClass
	return upClass, nil

}
