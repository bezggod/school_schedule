package office_usecase

import (
	"fmt"
	"school_schedule_2/internal/domain/model"
	"school_schedule_2/internal/domain/model/enums/office_name"
	"time"
)

type CreateOfficeReq struct {
	ID   int64
	Name office_name.OfficeName
}

func (u *OfficeUseCase) CreateOffice(req CreateOfficeReq) (*model.Office, error) {
	now := time.Now()
	office := &model.Office{req.ID, req.Name, now, now}
	err := u.r.CreateOffices(office)
	if err != nil {
		return nil, fmt.Errorf("u.r.CreateOffice: %w", err)
	}
	return office, nil
}
