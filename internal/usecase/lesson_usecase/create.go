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
	Office    enums.OfficeName
	TimeSlot  enums.TimeSlotName
	Subject   enums.SubjectName
	Date      time.Time
}

func (uc *UseCase) CreateLesson(req CreateLessonReq) (*model.Lesson, error) {
	lessonExist := uc.lessonRepo.LessonExists(req.Office, req.TimeSlot)

	if lessonExist {
		return nil, fmt.Errorf("lesson in office: %s, at time slot: %d already exist", req.Office, req.TimeSlot)
	}

	if _, err := uc.teacherRepo.GetByID(req.TeacherID); err != nil {
		return nil, fmt.Errorf("teacherRepo.GetByID: %w", err)
	}
	if _, err := uc.classRepo.GetByID(req.ClassID); err != nil {
		return nil, fmt.Errorf("classRepo.GetByID: %w", err)
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

	createdLesson, err := uc.lessonRepo.CreateLesson(lesson)
	if err != nil {
		return nil, fmt.Errorf("CreateLesson: %w", err)
	}

	return createdLesson, nil
}
