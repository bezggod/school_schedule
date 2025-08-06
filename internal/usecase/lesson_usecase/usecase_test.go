package lesson_usecase

import (
	"school_schedule_2/internal/usecase/lesson_usecase/mocks"
	"testing"
)

type mockService struct {
	lessonRepo  *mocks.LessonRepo
	teacherRepo *mocks.TeacherRepo
	classRepo   *mocks.ClassRepo
	timer       *mocks.Timer
}

func makeService(t *testing.T) (*UseCase, mockService) {
	m := mockService{
		lessonRepo:  mocks.NewLessonRepo(t),
		teacherRepo: mocks.NewTeacherRepo(t),
		classRepo:   mocks.NewClassRepo(t),
		timer:       mocks.NewTimer(t),
	}

	u := &UseCase{
		lessonRepo:  m.lessonRepo,
		teacherRepo: m.teacherRepo,
		classRepo:   m.classRepo,
		timer:       m.timer,
	}

	return u, m

}
