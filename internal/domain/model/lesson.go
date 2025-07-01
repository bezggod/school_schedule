package model

import (
	"fmt"
	"time"
)

type Lesson struct {
	ID        int64
	TeacherID int64
	ClassID   int64
	Office    *Office
	TimeSlot  *TimeSlot
	Subject   *Subject
	Date      time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewLesson(teacherID int64, classID int64, office *Office, timeSlot *TimeSlot, subject *Subject, date time.Time, now time.Time) *Lesson {
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
	office := "nil"
	if l.Office != nil {
		office = string(l.Office.Name)
	}

	timeSlot := "nil"
	if l.TimeSlot != nil {
		timeSlot = string(l.TimeSlot.Slot)
	}

	subject := "nil"
	if l.Subject != nil {
		subject = string(l.Subject.Name)
	}
	return fmt.Sprintf("Lesson[ID=%d, TeacherID=%d, ClassID=%d, Office=%s, TimeSlot=%s, Subject=%s, Date=%s]",
		l.ID,
		l.TeacherID,
		l.ClassID,
		office,
		timeSlot,
		subject,
		l.Date.Format("2006-01-02"),
	)

}
