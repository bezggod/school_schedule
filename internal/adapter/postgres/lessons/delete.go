package lessons

import (
	"context"
	"fmt"
)

func (r *Repo) DeleteLesson(id int64) error {
	_, err := r.cluster.Conn.Exec(context.Background(), `DELETE FROM lessons WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("r.cluster.Conn.Exec: %w", err)
	}
	return nil
}
