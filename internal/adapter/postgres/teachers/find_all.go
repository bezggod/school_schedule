package teachers

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"school_schedule_2/internal/domain/dto"
	"school_schedule_2/internal/domain/model"
)

func (r *Repo) FindAll(req dto.FindAllTeacherFilter) ([]*model.Teacher, error) {
	builder := squirrel.Select("id", "surname", "name", "patronymic", "created_at", "updated_at").From("teachers")
	if req.Surname != "" {
		builder = builder.Where("surname = ?", req.Surname)
	}

	if req.Name != "" {
		builder = builder.Where("name = ?", req.Name)
	}

	if req.Patronymic != "" {
		builder = builder.Where("patronymic = ?", req.Patronymic)
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

	teachers := make([]*model.Teacher, 0)
	for rows.Next() {
		var a model.Teacher
		err = rows.Scan(&a.ID, &a.Surname, &a.Name, &a.Patronymic, &a.CreatedAt, &a.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		teachers = append(teachers, &a)
		if err = rows.Err(); err != nil {
			return nil, err
		}
	}
	return teachers, nil
}
