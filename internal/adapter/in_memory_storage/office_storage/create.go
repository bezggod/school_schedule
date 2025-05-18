package office_storage

import "school_schedule_2/internal/domain/model"

func (r *OfficeRepo) CreateOffices(office *model.Office) error {
	office.ID = r.nextID
	r.offices[office.ID] = office
	r.nextID++

	return nil
}
