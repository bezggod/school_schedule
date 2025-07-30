package dto

type FindAllClassesFilter struct {
	ClassID int64
	Grade   string
}

type UpdateClassFilter struct {
	ClassID int64
	Grade   *string
}
