package lessons

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"school_schedule_2/internal/domain/model"
)

func (r *Repo) FindLesson(id int64) (*model.Lesson, error) {
	var lesson model.Lesson

	query := `
		SELECT id, teacher_id, class_id, office, time_slot, subject, date, created_at,update_at 
		FROM lessons 
		WHERE id = $1`

	err := r.cluster.Conn.QueryRow(context.Background(), query, id).
		Scan(
			&lesson.ID,
			&lesson.TeacherID,
			&lesson.ClassID,
			&lesson.Office,
			&lesson.TimeSlot,
			&lesson.Subject,
			&lesson.Date,
			&lesson.CreatedAt,
			&lesson.UpdatedAt,
		)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, NotFound
		}
		return nil, fmt.Errorf("FindLesson: %w", err)
	}
	return &lesson, nil
}
