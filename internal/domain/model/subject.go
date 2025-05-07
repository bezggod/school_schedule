package model

type Subject struct {
	ID   int8
	Name string
}

func NewSubjectMap() map[int8]Subject {
	return map[int8]Subject{
		1:  {ID: 1, Name: "История"},
		2:  {ID: 2, Name: "Обществознание"},
		3:  {ID: 3, Name: "Английский язык"},
		4:  {ID: 4, Name: "Физкультура"},
		5:  {ID: 5, Name: "Русский язык"},
		6:  {ID: 6, Name: "Литература"},
		7:  {ID: 7, Name: "Алгебра"},
		8:  {ID: 8, Name: "Геометрия"},
		9:  {ID: 9, Name: "Информатика"},
		10: {ID: 10, Name: "География"},
		11: {ID: 11, Name: "ОБЗР"},
	}
}
