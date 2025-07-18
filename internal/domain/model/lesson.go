package model

import (
	"fmt"
	"school_schedule_2/internal/domain/model/enums"
	"time"
)

type Lesson struct {
	ID        int64
	TeacherID int64
	ClassID   int64
	Office    enums.OfficeName
	TimeSlot  enums.TimeSlotName
	Subject   enums.SubjectName
	Date      time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewLesson(teacherID int64, classID int64, office enums.OfficeName, timeSlot enums.TimeSlotName, subject enums.SubjectName, date time.Time, now time.Time) *Lesson {
	return &Lesson{
		TeacherID: teacherID,
		ClassID:   classID,
		Office:    office,
		TimeSlot:  timeSlot,
		Subject:   subject,
		Date:      date,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
func (l *Lesson) String() string {
	return fmt.Sprintf("Lesson[ID=%d, TeacherID=%d, ClassID=%d, Office=%s, TimeSlot=%d, Subject=%s, Date=%s]",
		l.ID,
		l.TeacherID,
		l.ClassID,
		l.Office,
		l.TimeSlot,
		l.Subject,
		l.Date.Format("2006-01-02"),
	)

}
