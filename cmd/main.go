package main

import (
	"fmt"
	"school_schedule_2/cmd/service_provider"
	"school_schedule_2/internal/domain/model/enums"
	"school_schedule_2/internal/usecase/class_usecase"
	"school_schedule_2/internal/usecase/lesson_usecase"
	"school_schedule_2/internal/usecase/teacher_usecase"
	"time"
)

func main() {

	sp := service_provider.NewServiceProvider()

	createdTeacher, err := sp.GetTeacherUseCase().CreateTeacher(teacher_usecase.CreateTeacherReq{
		Surname:    "Безгубенко",
		Name:       "Данила",
		Patronymic: "Борисович",
	})
	if err != nil {
		fmt.Println("teacherUC.Create:", err)
		return
	}

	createdClass, err := sp.GetClassUseCase().CreateClass(class_usecase.CreateClassReq{
		Grade: "9М",
	})
	if err != nil {
		fmt.Println("classUC.Create:", err)
	}

	createdLesson, err := sp.GetLessonUseCase().CreateLesson(lesson_usecase.CreateLessonReq{
		TeacherID: createdTeacher.ID,
		ClassID:   createdClass.ID,
		Office:    enums.Office10,
		TimeSlot:  enums.FirstLesson,
		Subject:   enums.SubjectSocialScience,
		Date:      time.Now(),
	})
	if err != nil {
		fmt.Println("lessonUC.Create:", err)
		return
	}

	fmt.Printf("Lesson create: %+v\n", createdLesson)

	lessonsResp, err := sp.GetLessonUseCase().FindAll(lesson_usecase.FindAllReq{
		TeacherID: 1,
	})
	if err != nil {
		fmt.Println("lessonUC.FindAll:", err)
		return
	}

	classesResp, err := sp.GetClassUseCase().FindAll(class_usecase.FindAllReq{
		ClassID: 1,
	})
	if err != nil {
		fmt.Println("classUC.FindAll:", err)
		return
	}

	teacherReps, err := sp.GetTeacherUseCase().FindAll(teacher_usecase.FindAllReq{
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
