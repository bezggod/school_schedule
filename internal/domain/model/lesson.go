package model

import (
	"time"
)

type TimeSlot uint8

const (
	FirstLesson   TimeSlot = 1
	SecondLesson  TimeSlot = 2
	ThirdLesson   TimeSlot = 3
	FourthLesson  TimeSlot = 4
	FifthLesson   TimeSlot = 5
	SixthLesson   TimeSlot = 6
	SeventhLesson TimeSlot = 7
)

type Lesson struct {
	ID        int64
	Teacher   *Teacher
	Class     *Class
	Office    *Office
	TimeSlot  TimeSlot
	Subject   *Subject
	Date      time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewLesson(teacher *Teacher, class *Class, office *Office, timeSlot TimeSlot, subject *Subject, date time.Time, now time.Time) *Lesson {
	return &Lesson{
		Teacher:   teacher,
		Class:     class,
		Office:    office,
		TimeSlot:  timeSlot,
		Subject:   subject,
		Date:      date,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
