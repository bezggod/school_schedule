package classes

import (
	"context"
	"fmt"
	"school_schedule_2/internal/domain/model"
)

func (r *Repo) Create(class *model.Class) (*model.Class, error) {
	err := r.cluster.Conn.QueryRow(context.Background(),
		"INSERT INTO classes (grade, created_at, updated_at) VALUES ($1, $2, $3) "+
			"RETURNING id",
		class.Grade, class.CreatedAt, class.UpdatedAt).Scan(&class.ID)
	if err != nil {
		return nil, fmt.Errorf("Conn.QueryRow: %w", err)
	}
	return class, nil

}
