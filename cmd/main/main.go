package main

import (
	"fmt"
	"school_schedule_2/internal/adapter/in_memory_storage/class_storage"
	"school_schedule_2/internal/adapter/in_memory_storage/lesson_storage"
	"school_schedule_2/internal/adapter/in_memory_storage/teacher_storage"
	"school_schedule_2/internal/domain/model/enums"
	"school_schedule_2/internal/usecase/class_usecase"
	"school_schedule_2/internal/usecase/lesson_usecase"
	"school_schedule_2/internal/usecase/teacher_usecase"
	"time"
)

func main() {

	teacherRepo := teacher_storage.NewTeacherRepo()
	lessonRepo := lesson_storage.NewRepo()
	classRepo := class_storage.NewRepo()

	teacherUC := teacher_usecase.NewUseCase(teacherRepo)
	lessonUC := lesson_usecase.NewUseCase(lessonRepo, teacherRepo, classRepo)
	classUC := class_usecase.NewUseCase(classRepo)

	createdTeacher, err := teacherUC.CreateTeacher(teacher_usecase.CreateTeacherReq{
		Name:       "Безгубенко",
		Surname:    "Данила",
		Patronymic: "Борисович",
	})
	if err != nil {
		fmt.Println("teacherUC.CreateTeacher:", err)
		return
	}

	createdClass, err := classUC.CreateClass(class_usecase.CreateClassReq{
		Grade: "9М",
	})
	if err != nil {
		fmt.Println("classUC.CreateClass:", err)
	}

	createdLesson, err := lessonUC.CreateLesson(lesson_usecase.CreateLessonReq{
		TeacherID: createdTeacher.ID,
		ClassID:   createdClass.ID,
		Office:    enums.Office10,
		TimeSlot:  enums.FirstLesson,
		Subject:   enums.SubjectSocialScience,
		Date:      time.Now(),
	})
	if err != nil {
		fmt.Println("lessonUC.CreateLesson:", err)
		return
	}

	fmt.Printf("Lesson create: %+v\n", createdLesson)

	lessonsResp, err := lessonUC.FindAll(lesson_usecase.FindAllReq{
		TeacherID: 2,
	})
	if err != nil {
		fmt.Println("lessonUC.FindAll:", err)
		return
	}

	for _, lesson := range lessonsResp.Lessons {
		fmt.Printf("Lesson: %+v\n", lesson)
	}
}
