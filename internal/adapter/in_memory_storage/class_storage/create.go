package class_storage

import "school_schedule_2/internal/domain/model"

func (repo *ClassRepo) CreateClass(class *model.Class) (*model.Class, error) {
	class.ID = repo.nextID
	repo.nextID++
	repo.classes[class.ID] = class
	return class, nil
}
