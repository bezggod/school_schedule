package lesson_usecase

import (
	"errors"
	"school_schedule_2/internal/domain/model"
	"time"
)

type CreateLessonReq struct {
	Teacher  *model.Teacher
	Class    *model.Class
	Office   *model.Office
	TimeSlot model.TimeSlot
	Subject  *model.Subject
	Date     time.Time
}

func (usecase *LessonUseCase) CreateLesson(req CreateLessonReq) (*model.Lesson, error) {
	if req.Teacher == nil || req.Class == nil || req.Office == nil {
		return nil, errors.New("Teacher, Class и Office не могут быть nil")
	}

	now := time.Now()
	lesson := model.NewLesson(req.Teacher, req.Class, req.Office, req.TimeSlot, req.Subject, req.Date, now)

	createdLesson, err := usecase.LessonRepo.CreateLesson(lesson)
	if err != nil {
		return nil, err
	}

	return createdLesson, nil
}
