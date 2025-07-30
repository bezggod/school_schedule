package dto

type FindAllTeacherFilter struct {
	Surname    string
	Name       string
	Patronymic string
}

type UpdateTeacherFilter struct {
	ID         int64
	Surname    *string
	Name       *string
	Patronymic *string
}
