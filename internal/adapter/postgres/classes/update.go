package classes

import (
	"context"
	"fmt"
	"school_schedule_2/internal/domain/model"
	"time"
)

func (r *Repo) Update(class *model.Class) (*model.Class, error) {
	class.UpdatedAt = time.Now()

	_, err := r.cluster.Conn.Exec(context.Background(),
		`UPDATE classes SET grade = $1, updated_at = $2 WHERE id = $3`,
		class.Grade, class.UpdatedAt, class.ID)

	if err != nil {
		return nil, fmt.Errorf("Update.Exec: %w", err)
	}

	return class, nil
}
