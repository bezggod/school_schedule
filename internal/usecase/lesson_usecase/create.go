package lesson_usecase

import (
	"errors"
	"fmt"
	"time"

	"school_schedule_2/internal/domain/model"
	"school_schedule_2/internal/domain/model/enums"
)

var errLessonAlreadyExists = errors.New("lesson already exists")

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
		return nil, errors.Join(errLessonAlreadyExists, fmt.Errorf("lesson in office: %s, at time slot: %d", req.Office, req.TimeSlot))
	}

	if _, err := uc.teacherRepo.GetByID(req.TeacherID); err != nil {
		return nil, fmt.Errorf("teacherRepo.GetByID: %w", err)
	}

	// Find - вернет nil если класса нет в базе данных
	// Get - вернет ошибку если нет класса в базе данных

	_, err := uc.classRepo.GetByID(req.ClassID)
	if err != nil {
		return nil, fmt.Errorf("classRepo.GetByID: %w", err)
	}

	now := uc.timer.NowMoscow()
	lesson := model.NewLesson(
		req.TeacherID,
		req.ClassID,
		req.Office,
		req.TimeSlot,
		req.Subject,
		req.Date,
		now)

	createdLesson, err := uc.lessonRepo.Create(lesson)
	if err != nil {
		return nil, fmt.Errorf("Create: %w", err)
	}

	return createdLesson, nil
}
