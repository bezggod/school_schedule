package lesson_usecase

import (
	"fmt"
	"school_schedule_2/internal/domain/model"
	"school_schedule_2/internal/domain/model/enums"
)

type UpdateLessonReq struct {
	ID        int64
	TeacherID int64
	ClassID   int64
	Office    enums.OfficeName
	TimeSlot  enums.TimeSlotName
	Subject   enums.SubjectName
}

func (ul *UseCase) Update(req UpdateLessonReq) (*model.Lesson, error) {
	lesson, err := ul.lessonRepo.GetByID(req.ID)
	if err != nil {
		return nil, fmt.Errorf("lessonRepo.GetByID: %w", err)
	}

	lesson.TeacherID = req.TeacherID
	lesson.ClassID = req.ClassID
	lesson.Office = req.Office
	lesson.TimeSlot = req.TimeSlot
	lesson.Subject = req.Subject

	lesson, err = ul.lessonRepo.Update(lesson)
	if err != nil {
		return nil, fmt.Errorf("lessonRepo.Update: %w", err)
	}

	return lesson, nil
}
