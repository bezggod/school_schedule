package classes

import (
	"context"
	"fmt"
	"school_schedule_2/internal/domain/model"
)

func (r *Repo) UpdateClass(class *model.Class) (*model.Class, error) {
	query := `
		UPDATE classes SET grade =$1, updated_at = $2 
        WHERE id = $3`

	_, err := r.cluster.Conn.Exec(context.Background(), query,
		class.Grade,
		class.UpdatedAt,
		class.ID,
	)
	if err != nil {
		return nil, fmt.Errorf("UpdateClass: failed to update class: %w", err)
	}
	return class, nil
}
