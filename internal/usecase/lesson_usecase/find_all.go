package lesson_usecase

import (
	"errors"
	"fmt"
	"school_schedule_2/internal/domain/dto"
	"school_schedule_2/internal/domain/model"
	"school_schedule_2/internal/domain/model/enums"
	"time"
)

var errEmptyTeacher = errors.New("empty teacher name")

type FindAllReq struct {
	TeacherID int64
	ClassID   int64
	Office    enums.OfficeName
	TimeSlot  enums.TimeSlotName
	Subject   enums.SubjectName
}
type FindAllResponse struct {
	Lessons []*Lesson
}

type Lesson struct {
	ID        int64
	Teacher   *model.Teacher
	Class     *model.Class
	Office    enums.OfficeName
	TimeSlot  enums.TimeSlotName
	Subject   enums.SubjectName
	Date      time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (uc *UseCase) FindAll(req FindAllReq) (*FindAllResponse, error) {
	lessons, err := uc.lessonRepo.FindLesson(dto.FindAllLessonFilter{
		TeacherID:  req.TeacherID,
		OfficeName: req.Office,
		TimeSlot:   req.TimeSlot,
		ClassID:    req.ClassID,
	})
	if err != nil {
		return nil, fmt.Errorf("lessonRepo.FindAll: %w", err)
	}

	resultLessons := make([]*Lesson, 0, len(lessons))
	for _, lesson := range lessons {

		var (
			teacher *model.Teacher
			class   *model.Class
		)

		teacher, err = uc.teacherRepo.GetByID(lesson.TeacherID)
		if err != nil {
			return nil, fmt.Errorf("teacherRepo.GetByID: %w", err)
		}
		if teacher == nil {
			return nil, fmt.Errorf("teacherRepo.GetByID: %w with id %d", errEmptyTeacher, lesson.TeacherID)
		}

		class, err = uc.classRepo.GetByID(lesson.ClassID)
		if err != nil {
			return nil, fmt.Errorf("classRepo.GetByID: %w", err)
		}

		resultLessons = append(resultLessons, &Lesson{ //Отражение данных на фронте
			ID:        lesson.ID,
			Teacher:   teacher,
			Class:     class,
			Office:    lesson.Office,
			TimeSlot:  lesson.TimeSlot,
			Subject:   lesson.Subject,
			Date:      lesson.Date,
			CreatedAt: lesson.CreatedAt,
			UpdatedAt: lesson.UpdatedAt,
		})
	}

	return &FindAllResponse{Lessons: resultLessons}, nil
}
