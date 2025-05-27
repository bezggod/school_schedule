package office_usecase

import "school_schedule_2/internal/adapter/in_memory_storage/office_storage"

type OfficeUseCase struct {
	r *office_storage.OfficeRepo
}

func NewOfficeUseCase(r *office_storage.OfficeRepo) *OfficeUseCase {
	return &OfficeUseCase{r}
}
