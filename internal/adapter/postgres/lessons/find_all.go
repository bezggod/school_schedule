package lessons

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"school_schedule_2/internal/domain/dto"
	"school_schedule_2/internal/domain/model"
)

func (r *Repo) FindAll(req dto.FindAllLessonFilter) ([]*model.Lesson, error) {
	builder := squirrel.Select(
		"teacherID",
		"classID",
		"office",
		"timeSlot",
		"subject",
	).
		From("lessons")
	if req.TeacherID != 0 {
		builder = builder.Where("teacherID = ?", req.TeacherID)
	}
	if req.ClassID != 0 {
		builder = builder.Where("classID = ?", req.ClassID)
	}
	if req.OfficeName != "" {
		builder = builder.Where("officeName = ?", req.OfficeName)
	}
	if req.TimeSlot != 0 {
		builder = builder.Where("timeSlot = ?", req.TimeSlot)
	}
	if req.Subject != "" {
		builder = builder.Where("subject = ?", req.Subject)
	}
	query, _, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	rows, err := r.cluster.Conn.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("failed to execiute query: %w", err)
	}
	defer rows.Close()

	lessons := make([]*model.Lesson, 0)
	for rows.Next() {
		var l model.Lesson
		err = rows.Scan(&l.TeacherID, &l.ClassID, &l.Office, &l.TimeSlot, &l.Subject)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		lessons = append(lessons, &l)
		if err = rows.Err(); err != nil {
			return nil, err
		}
	}
	return lessons, nil
}
