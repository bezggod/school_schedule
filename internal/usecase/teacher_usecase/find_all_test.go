package teacher_usecase

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"school_schedule_2/internal/domain/dto"
	"school_schedule_2/internal/domain/model"
	"testing"
	"time"
)

func TestUseCase_FindAll(t *testing.T) {
	t.Parallel()

	now := time.Now()
	errTest := errors.New("test error")

	type args struct {
		req FindAllReq
	}

	tests := []struct {
		name    string
		args    args
		want    []*model.Teacher
		wantErr error
		before  func(m mockService, args args)
	}{
		{
			name: "success",
			args: args{
				req: FindAllReq{
					Surname: "Безгубенко",
				},
			},
			want: []*model.Teacher{
				{
					Surname:    "Безгубенко",
					Name:       "Данила",
					Patronymic: "Борисович",
					CreatedAt:  now,
					UpdatedAt:  now,
				},
			},
			before: func(m mockService, args args) {
				teachers := []*model.Teacher{
					{
						Surname:    "Безгубенко",
						Name:       "Данила",
						Patronymic: "Борисович",
						CreatedAt:  now,
						UpdatedAt:  now,
					},
				}
				m.teacherRepo.EXPECT().FindAll(dto.FindAllTeacherFilter{
					Surname:    args.req.Surname,
					Name:       args.req.Name,
					Patronymic: args.req.Patronymic,
				}).Return(teachers, nil)
			},
		},
		{
			name: "teacherRepo.FindAll err",
			args: args{
				req: FindAllReq{
					Surname:    "",
					Name:       "",
					Patronymic: "",
				},
			},
			wantErr: errTest,
			before: func(m mockService, args args) {
				m.teacherRepo.EXPECT().FindAll(dto.FindAllTeacherFilter{
					Surname:    args.req.Surname,
					Name:       args.req.Name,
					Patronymic: args.req.Patronymic,
				}).Return(nil, errTest)
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

			teachers, err := usecase.FindAll(tc.args.req)
			if tc.wantErr != nil {
				a.Error(err)
				a.Contains(err.Error(), tc.wantErr.Error())

				return
			}
			a.NoError(err)
			a.Equal(tc.want, teachers)
		})
	}
}
