package main

import (
	"fmt"
	"school_schedule_2/internal/adapter/in_memory_storage/class_storage"
	"school_schedule_2/internal/adapter/in_memory_storage/lesson_storage"
	"school_schedule_2/internal/adapter/in_memory_storage/teacher_storage"
	"school_schedule_2/internal/domain/model"
	"school_schedule_2/internal/domain/model/enums"
	"school_schedule_2/internal/usecase/class_usecase"
	"school_schedule_2/internal/usecase/lesson_usecase"
	"school_schedule_2/internal/usecase/teacher_usecase"
	"time"
)

func main() {

	teacherRepo := teacher_storage.NewTeacherRepo()
	lessonRepo := lesson_storage.NewClassRepo()
	classRepo := class_storage.NewClassRepo()

	teacherUC := teacher_usecase.NewTeacherUseCase(teacherRepo)
	lessonUC := lesson_usecase.NewLessonUseCase(lessonRepo, teacherRepo)
	classUC := class_usecase.NewUseCase(classRepo)

	createdTeacher, err := teacherUC.CreateTeacher(teacher_usecase.CreateTeacherReq{
		Name:       "Безгубенко",
		Surname:    "Данила",
		Patronymic: "Борисович",
	})
	if err != nil {
		fmt.Println("Ошибка в создании учителя", err)
		return
	}

	createdClass, err := classUC.CreateClass(class_usecase.CreateClassReq{
		Grade: "9М",
	})
	if err != nil {
		fmt.Println("Ошибка создания класса", err)
	}

	timeslot := enums.FirstLesson
	office := enums.Office10
	subject := enums.SubjectSocialScience

	createdLesson, err := lessonUC.CreateLesson(lesson_usecase.CreateLessonReq{
		TeacherID: createdTeacher.ID,
		ClassID:   createdClass.ID,
		Office:    model.Office,
		TimeSlot:  timeslot,
		Subject:   subject,
		Date:      time.Now(),
	})
	if err != nil {
		fmt.Println("Ошибка создания урока:", err)
		return
	}

	fmt.Printf("Создан урок: %+v\n", createdLesson)
}
