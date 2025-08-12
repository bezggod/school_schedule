package lessons

import (
	"context"
	"school_schedule_2/internal/domain/model/enums"
)

func (r *Repo) LessonExists(name enums.OfficeName, slot enums.TimeSlotName) bool {
	query := `SELECT 1 FROM lessons WHERE office = $1 AND time_slot = $2 LIMIT 1`

	var exists int
	err := r.cluster.Conn.QueryRow(context.Background(), query, name, slot).Scan(&exists)
	return err == nil
}
