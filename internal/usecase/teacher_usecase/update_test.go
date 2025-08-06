package teacher_usecase

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"school_schedule_2/internal/domain/model"
	"testing"
	"time"
)

func TestUseCase_UpdateTeacher(t *testing.T) {
	t.Parallel()

	now := time.Now()
	errTest := errors.New("update test error")

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
					Surname:    "Безгубенко",
					Name:       "Данила",
					Patronymic: "Борисович",
				},
			},
			want: &model.Teacher{
				ID:         1,
				Surname:    "Безгубенко",
				Name:       "Данила",
				Patronymic: "Борисович",
				CreatedAt:  now,
				UpdatedAt:  now,
			},
			before: func(m mockService, args args) {
				teacher := &model.Teacher{
					ID:         1,
					Surname:    "Безгубенко",
					Name:       "Данила",
					Patronymic: "Борисович",
					CreatedAt:  now,
					UpdatedAt:  now,
				}

				m.teacherRepo.EXPECT().GetByID(args.req.ID).Return(teacher, nil)

				teacher.Surname = args.req.Surname
				teacher.Name = args.req.Name
				teacher.Patronymic = args.req.Patronymic

				m.teacherRepo.EXPECT().Update(teacher).Return(teacher, nil)
			},
		},
		{
			name: "error on teacher get by id",
			args: args{
				req: UpdateTeacherReq{
					ID:         1,
					Surname:    "Безгубенко",
					Name:       "Данила",
					Patronymic: "Борисович",
				},
			},
			wantErr: errTest,
			before: func(m mockService, args args) {

				m.teacherRepo.EXPECT().GetByID(args.req.ID).Return(nil, errTest)
			},
		},
		{
			name: "error on teacher update",
			args: args{
				req: UpdateTeacherReq{
					ID:         1,
					Surname:    "Безгубенко",
					Name:       "Данила",
					Patronymic: "Борисович",
				},
			},
			wantErr: errTest,
			before: func(m mockService, args args) {
				teacher := &model.Teacher{
					ID:         1,
					Surname:    "Безгубенко",
					Name:       "Данила",
					Patronymic: "Борисович",
					CreatedAt:  now,
					UpdatedAt:  now,
				}

				m.teacherRepo.EXPECT().GetByID(args.req.ID).Return(teacher, nil)

				m.teacherRepo.EXPECT().Update(teacher).Return(nil, errTest)
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

			teachers, err := usecase.Update(tc.args.req)
			if tc.wantErr != nil {
				a.Error(err)
				a.Contains(err.Error(), tc.wantErr.Error())
				return
			}
			a.Equal(tc.want, teachers)
			a.NoError(err)
		})
	}
}
