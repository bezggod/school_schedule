package main

import (
	"fmt"
	"school_schedule_2/internal/adapter/in_memory_storage/class_storage"
	"school_schedule_2/internal/adapter/in_memory_storage/lesson_storage"
	"school_schedule_2/internal/adapter/in_memory_storage/teacher_storage"
	"school_schedule_2/internal/domain/model/enums"
	"school_schedule_2/internal/pkg"
	"school_schedule_2/internal/usecase/class_usecase"
	"school_schedule_2/internal/usecase/lesson_usecase"
	"school_schedule_2/internal/usecase/teacher_usecase"
	"time"
)

func main() {

	teacherRepo := teacher_storage.NewRepo()
	lessonRepo := lesson_storage.NewRepo()
	classRepo := class_storage.NewRepo()
	timer := pkg.NewTimer()

	teacherUC := teacher_usecase.NewUseCase(teacherRepo, timer)
	lessonUC := lesson_usecase.NewUseCase(lessonRepo, teacherRepo, classRepo)
	classUC := class_usecase.NewUseCase(classRepo)

	createdTeacher, err := teacherUC.CreateTeacher(teacher_usecase.CreateTeacherReq{
		Surname:    "Безгубенко",
		Name:       "Данила",
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
		TeacherID: 1,
	})
	if err != nil {
		fmt.Println("lessonUC.FindAll:", err)
		return
	}

	classesResp, err := classUC.FindAll(class_usecase.FindAllReq{
		ClassID: 1,
	})
	if err != nil {
		fmt.Println("classUC.FindAll:", err)
		return
	}

	teacherReps, err := teacherUC.FindAll(teacher_usecase.FindAllReq{
		Surname: "Безгубенко",
	})
	if err != nil {
		fmt.Println("teacherUC.FindAll:", err)
	}

	for _, lesson := range lessonsResp.Lessons {
		fmt.Printf("Lesson: %+v\n", lesson)
	}

	for _, class := range classesResp {
		fmt.Printf("Class: %+v\n", class)
	}

	for _, teacher := range teacherReps {
		fmt.Printf("Teacher: %+v\n", teacher)
	}

}
