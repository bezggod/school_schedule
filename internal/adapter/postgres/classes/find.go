package classes

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"school_schedule_2/internal/domain/model"
)

func (r *Repo) FindClass(id int64) (*model.Class, error) {
	var class model.Class

	query := `
		SELECT id, grade ,created_at, updated_at 
		FROM classes 
		WHERE id = $1`

	err := r.cluster.Conn.QueryRow(context.Background(), query, id).
		Scan(
			&class.ID,
			&class.Grade,
			&class.CreatedAt,
			&class.UpdatedAt,
		)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("no class found with id %d", id)
		}
		return nil, fmt.Errorf("FindClass: %w", err)
	}
	return &class, nil

}
