package teacher_usecase

import (
	"school_schedule_2/internal/usecase/teacher_usecase/mocks"
	"testing"
)

type mockService struct {
	teacherRepo *mocks.TeacherRepo
	timer       *mocks.Timer
}

func makeService(t *testing.T) (*UseCase, mockService) {
	m := mockService{
		teacherRepo: mocks.NewTeacherRepo(t),
		timer:       mocks.NewTimer(t),
	}

	u := &UseCase{
		teacherRepo: m.teacherRepo,
		timer:       m.timer,
	}

	return u, m

}
