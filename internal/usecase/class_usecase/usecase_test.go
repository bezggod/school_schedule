package class_usecase

import (
	"school_schedule_2/internal/usecase/class_usecase/mocks"
	"testing"
)

type mockService struct {
	classRepo *mocks.ClassRepo
	timer     *mocks.Timer
}

func makeService(t *testing.T) (*UseCase, mockService) {
	m := mockService{
		classRepo: mocks.NewClassRepo(t),
		timer:     mocks.NewTimer(t),
	}

	u := &UseCase{
		classRepo: m.classRepo,
		timer:     m.timer,
	}

	return u, m

}
