package classes

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"school_schedule_2/internal/domain/dto"
	"school_schedule_2/internal/domain/model"
)

func (r *Repo) FindAll(req dto.FindAllClassesFilter) ([]*model.Class, error) {
	builder := squirrel.Select("class_id", "grade", "created_at", "updated_at").From("classes")
	if req.ClassID != 0 {
		builder = builder.Where("class_id = ?", req.ClassID)
	}

	if req.Grade != "" {
		builder = builder.Where("grade = ?", req.Grade)
	}

	query, _, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	rows, err := r.cluster.Conn.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	classes := make([]*model.Class, 0)
	for rows.Next() {
		var a model.Class
		err = rows.Scan(&a.ID, &a.Grade, &a.CreatedAt, &a.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		classes = append(classes, &a)
		if err = rows.Err(); err != nil {
			return nil, err
		}
	}
	return classes, nil
}
