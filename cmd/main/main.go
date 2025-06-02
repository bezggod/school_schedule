package main

import (
	"fmt"
	"school_schedule_2/internal/domain/model/enums"
	"time"

	"school_schedule_2/internal/adapter/in_memory_storage/lesson_storage"
	"school_schedule_2/internal/adapter/in_memory_storage/teacher_storage"
	"school_schedule_2/internal/domain/model"
	"school_schedule_2/internal/usecase/lesson_usecase"
)

func main() {
	teacherRepo := teacher_storage.NewTeacherRepo()
	lessonRepo := lesson_storage.NewLessonRepo()

	now := time.Now()

	teacher := model.NewTeacher("Безгубенко", "Данила", "Борисович", now)

	createdTeacher, err := teacherRepo.CreateTeacher(teacher)
	if err != nil {
		fmt.Println("Ошибка при создании учителя:", err)
		return
	}

	teacher.ID = createdTeacher.ID

	office := &model.Office{
		ID:   1,
		Name: enums.Office10, // Константа из enums
	}

	class := &model.Class{
		ID:    1,
		Grade: "9М",
	}

	lessonUC := lesson_usecase.NewLessonUseCase(lessonRepo)

	req := lesson_usecase.CreateLessonReq{
		Teacher:  teacher,
		Class:    class,
		Office:   office,
		TimeSlot: model.FirstLesson,
		Date:     now,
	}

	lesson, err := lessonUC.CreateLesson(req)
	if err != nil {
		fmt.Println("Ошибка создания урока:", err)
		return
	}

	fmt.Printf("Создан урок: %+v\n", lesson)
}
