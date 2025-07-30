package teacher_usecase

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"school_schedule_2/internal/domain/model"
	"testing"
	"time"
)

func TestUseCase_UpdateTeacher(t *testing.T) {
	t.Parallel()

	now := time.Now()
	errTest := errors.New("update test error")

	surname := "surname"
	name := "name"
	patronymic := "patronymic"

	type args struct {
		req UpdateTeacherReq
	}

	tests := []struct {
		name    string
		args    args
		want    *model.Teacher
		wantErr error
		before  func(m mockService, args args)
	}{
		{
			name: "success",
			args: args{
				req: UpdateTeacherReq{
					ID:         1,
					Surname:    &surname,
					Name:       &name,
					Patronymic: &patronymic,
				},
			},
			want: &model.Teacher{
				ID:         1,
				Surname:    surname,
				Name:       name,
				Patronymic: patronymic,
				CreatedAt:  now,
				UpdatedAt:  now,
			},
			before: func(m mockService, args args) {
				teacherOld := &model.Teacher{
					ID:         1,
					Surname:    "Old",
					Name:       "Old",
					Patronymic: "Old",
					CreatedAt:  now,
					UpdatedAt:  now,
				}
				teacherNew := &model.Teacher{
					ID:         1,
					Surname:    *args.req.Surname,
					Name:       *args.req.Name,
					Patronymic: *args.req.Patronymic,
					CreatedAt:  now,
					UpdatedAt:  now,
				}
				m.teacherRepo.EXPECT().GetByID(args.req.ID).Return(teacherOld, nil)
				m.teacherRepo.EXPECT().UpdateTeacher(teacherNew).Return(teacherNew, nil)
			},
		},
		{
			name: "invalid teacher id",
			args: args{
				req: UpdateTeacherReq{
					ID: 0,
				},
			},
			wantErr: fmt.Errorf("invalid teacher id"),
		},
		{
			name: "teacherRepo.GetByID error",
			args: args{
				req: UpdateTeacherReq{
					ID: 1,
				},
			},
			wantErr: errTest,
			before: func(m mockService, args args) {
				m.teacherRepo.EXPECT().GetByID(args.req.ID).Return(nil, errTest)
			},
		},
		{
			name: "teacherRepo.UpdateTeacher error",
			args: args{
				req: UpdateTeacherReq{
					ID:      1,
					Surname: &surname,
				},
			},
			wantErr: errTest,
			before: func(m mockService, args args) {
				teacherOld := &model.Teacher{
					ID:         1,
					Surname:    "",
					Name:       "",
					Patronymic: "",
				}
				teacherNew := &model.Teacher{
					ID:         1,
					Surname:    *args.req.Surname,
					Name:       "",
					Patronymic: "",
				}
				m.teacherRepo.EXPECT().GetByID(args.req.ID).Return(teacherOld, nil)
				m.teacherRepo.EXPECT().UpdateTeacher(teacherNew).Return(nil, errTest)
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

			teachers, err := usecase.UpdateTeacher(tc.args.req)
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
