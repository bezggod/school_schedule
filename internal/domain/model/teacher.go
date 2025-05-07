package model

import "time"

type Teacher struct {
	ID        int64
	Name      string
	Surname   string
	CreatedAt time.Time
	UpdatedAt time.Time
	Subject   map[int8]Subject
}

//func NewTeacher(Name, Surname string) *Teacher {
//
//}
