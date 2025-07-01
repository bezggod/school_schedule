package lesson_usecase

import (
	"fmt"
	"school_schedule_2/internal/domain/model"
	"school_schedule_2/internal/domain/model/enums"
	"time"
)

type CreateLessonReq struct {
	TeacherID int64
	ClassID   int64
	Office    *model.Office
	TimeSlot  *model.TimeSlot
	Subject   *model.Subject
	Date      time.Time
}

func (u *LessonUseCase) CreateLesson(req CreateLessonReq) (*model.Lesson, error) {
	lessonExist := u.lessonRepo.LessonExists(
		enums.Office10,
		model.TimeSlot{Slot: enums.FirstLesson})

	if lessonExist {
		return nil, fmt.Errorf("%d", enums.Office10, enums.FirstLesson)
	}

	if _, err := u.teacherRepo.GetByID(req.TeacherID); err != nil {
		return nil, err
	}
	if _, err := u.classRepo.GetByID(req.ClassID); err != nil {
		return nil, err
	}

	now := time.Now()
	lesson := model.NewLesson(
		req.TeacherID,
		req.ClassID,
		req.Office,
		req.TimeSlot,
		req.Subject,
		req.Date,
		now)

	createdLesson, err := u.lessonRepo.CreateLesson(lesson)
	if err != nil {
		return nil, err
	}

	return createdLesson, nil
}
