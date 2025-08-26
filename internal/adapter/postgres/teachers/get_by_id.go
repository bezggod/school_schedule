package teachers

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"school_schedule_2/internal/domain/model"
)

func (r *Repo) GetByID(id int64) (*model.Teacher, error) {
	var teacher model.Teacher

	query := `
		SELECT id,surname, name, patronymic, created_at, updated_at 
		FROM teachers 	
		WHERE id = $1`

	err := r.cluster.Conn.QueryRow(context.Background(), query, id).
		Scan(
			&teacher.ID,
			&teacher.Surname,
			&teacher.Name,
			&teacher.Patronymic,
			&teacher.CreatedAt,
			&teacher.UpdatedAt,
		)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, NotFound
		}
		return nil, fmt.Errorf("GetByID.Teacher: %w", err)
	}
	return &teacher, nil
}
