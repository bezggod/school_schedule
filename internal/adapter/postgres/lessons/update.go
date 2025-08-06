package lessons

import (
	"context"
	"fmt"
	"school_schedule_2/internal/domain/model"
)

func (r *Repo) UpdateLesson(lesson *model.Lesson) (*model.Lesson, error) {
	query := `
		UPDATE lessons SET teacher_id=$1, class_id=$2, office=$3,time_slot=$4, subject=$5,updated_at=6 
        WHERE id=$7;`

	_, err := r.cluster.Conn.Exec(context.Background(), query,
		lesson.TeacherID,
		lesson.ClassID,
		lesson.Office,
		lesson.TimeSlot,
		lesson.Subject,
		lesson.UpdatedAt,
		lesson.ID,
	)
	if err != nil {
		return nil, fmt.Errorf("UpdateLesson: failed to update lesson %w", err)
	}
	return lesson, nil
}
