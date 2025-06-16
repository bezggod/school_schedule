package main

import (
	"fmt"
	"school_schedule_2/internal/adapter/in_memory_storage/class_storage"
	"school_schedule_2/internal/domain/model/enums"
	"school_schedule_2/internal/usecase/class_usecase"
	"school_schedule_2/internal/usecase/teacher_usecase"
	"time"

	"school_schedule_2/internal/adapter/in_memory_storage/lesson_storage"
	"school_schedule_2/internal/adapter/in_memory_storage/teacher_storage"
	"school_schedule_2/internal/domain/model"
	"school_schedule_2/internal/usecase/lesson_usecase"
)

func main() {
	teacherRepo := teacher_storage.NewTeacherRepo()
	lessonRepo := lesson_storage.NewLessonRepo()
	classRepo := class_storage.NewClassRepo()

	teacherUC := teacher_usecase.NewTeacherUseCase(teacherRepo)
	lessonUC := lesson_usecase.NewLessonUseCase(lessonRepo)
	classUC := class_usecase.NewClassUseCase(classRepo)

	teacherReq := teacher_usecase.CreateTeacherReq{
		Name:       "Безгубенко",
		Surname:    "Данила",
		Patronymic: "Борисович",
	}

	createdTeacher, err := teacherUC.CreateTeacher(teacherReq)
	if err != nil {
		fmt.Println("Ошибка в создании учителя", err)
		return
	}

	classReq := class_usecase.CreateClassReq{
		Grade: "9М",
	}
	createdClass, err := classUC.CreateClass(classReq)
	if err != nil {
		fmt.Println("Ошибка создания класса", err)
	}

	office := model.NewOffice(enums.Office10, time.Now())
	subject := model.NewSubject(enums.SubjectHistory, time.Now())

	lessonReq := lesson_usecase.CreateLessonReq{
		Teacher:  createdTeacher,
		Class:    createdClass,
		Office:   office,
		TimeSlot: model.FirstLesson,
		Subject:  subject,
		Date:     time.Now(),
	}

	createdLesson, err := lessonUC.CreateLesson(lessonReq)
	if err != nil {
		fmt.Println("Ошибка создания урока:", err)
		return
	}

	fmt.Printf("Создан урок: %+v\n", createdLesson)
}
