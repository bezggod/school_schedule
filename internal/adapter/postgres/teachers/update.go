package teachers

import (
	"context"
	"fmt"
	"school_schedule_2/internal/domain/model"
)

func (r *Repo) UpdateTeacher(teacher *model.Teacher) (*model.Teacher, error) {
	query := `
		UPDATE teachers SET surname =$1,name=$2,patronymic=$3, updated_at = $4 
        WHERE id = $5`

	_, err := r.cluster.Conn.Exec(context.Background(), query,
		teacher.Surname,
		teacher.Name,
		teacher.Patronymic,
		teacher.UpdatedAt,
		teacher.ID,
	)
	if err != nil {
		return nil, fmt.Errorf("UpdateTeacher: failed to update teacher: %w", err)
	}
	return teacher, nil
}
