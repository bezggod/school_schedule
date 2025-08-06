package teachers

import (
	"context"
	"fmt"
	"school_schedule_2/internal/domain/model"
)

func (r *Repo) CreateTeacher(teacher *model.Teacher) (*model.Teacher, error) {
	err := r.cluster.Conn.QueryRow(context.Background(),
		"INSERT INTO teachers(surname,name,patronymic,created_at,updated_at) VALUES($1,$2,$3,$4,$5) "+
			"returning id",
		teacher.Surname, teacher.Name, teacher.Patronymic, teacher.CreatedAt, teacher.UpdatedAt).Scan(&teacher.ID)
	if err != nil {
		return nil, fmt.Errorf("Conn.QueryRow err: %w", err)
	}
	return teacher, nil
}
