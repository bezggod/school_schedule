package teacher_usecase

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"school_schedule_2/internal/domain/model"
	"testing"
	"time"
)

func TestUseCase_CreateTeacher(t *testing.T) {
	t.Parallel()

	now := time.Now()
	errTest := errors.New("test error")

	type args struct {
		req CreateTeacherReq
	}

	tests := []struct {
		name    string
		args    args
		want    *model.Teacher
		wantErr bool
		before  func(m mockService, args args)
	}{
		{
			name: "success",
			args: args{
				req: CreateTeacherReq{
					Surname:    "Безгубенко",
					Name:       "Данила",
					Patronymic: "Борисович",
				},
			},
			want: &model.Teacher{
				Surname:    "Безгубенко",
				Name:       "Данила",
				Patronymic: "Борисович",
			},
			before: func(m mockService, args args) {
				m.timer.EXPECT().NowMoscow().Return(now)

				teacher := &model.Teacher{
					Surname:    args.req.Surname,
					Name:       args.req.Name,
					Patronymic: args.req.Patronymic,
					CreatedAt:  now,
					UpdatedAt:  now,
				}

				m.teacherRepo.EXPECT().Create(teacher).Return(&model.Teacher{
					Surname:    "Безгубенко",
					Name:       "Данила",
					Patronymic: "Борисович",
				}, nil)
			},
		},
		{
			name: "error on Create",
			args: args{
				req: CreateTeacherReq{
					Surname:    "Безгубенко",
					Name:       "Данила",
					Patronymic: "Борисович",
				},
			},
			wantErr: true,
			before: func(m mockService, args args) {
				m.timer.EXPECT().NowMoscow().Return(now)

				teacher := &model.Teacher{
					Surname:    args.req.Surname,
					Name:       args.req.Name,
					Patronymic: args.req.Patronymic,
					CreatedAt:  now,
					UpdatedAt:  now,
				}
				m.teacherRepo.EXPECT().Create(teacher).Return(nil, errTest)
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			a := assert.New(t)

			usecase, mocks := makeService(t)
			if tc.before != nil {
				tc.before(mocks, tc.args)
			}

			teacher, err := usecase.CreateTeacher(tc.args.req)
			if tc.wantErr {
				a.Error(err)

				return
			}
			a.NoError(err)
			a.Equal(tc.want, teacher)
		})
	}
}
