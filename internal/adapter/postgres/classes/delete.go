package classes

import (
	"context"
	"fmt"
)

func (r *Repo) DeleteClass(id int) error {
	_, err := r.cluster.Conn.Exec(context.Background(),
		`DELETE FROM classes WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("r.cluster.Conn.Exec: %w", err)
	}
	return nil
}
