package lessons

import (
	"context"
	"fmt"
	"school_schedule_2/internal/domain/model"
)

func (r *Repo) Create(lesson *model.Lesson) (*model.Lesson, error) {
	err := r.cluster.Conn.QueryRow(context.Background(),
		"INSERT INTO lessons (teacher_id, class_id, office, time_slot,subject, date, created_at,updated_at) values ($1,$2,$3,$4,$5,$6,$7,$8)"+
			"RETURNING id",
		lesson.TeacherID, lesson.ClassID, lesson.Office, lesson.TimeSlot, lesson.Subject, lesson.Date, lesson.CreatedAt, lesson.UpdatedAt).Scan(&lesson.ID)
	if err != nil {
		return nil, fmt.Errorf("Conn.QueryRow: %w", err)
	}
	return lesson, nil
}
