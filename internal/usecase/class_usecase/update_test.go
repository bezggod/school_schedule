package class_usecase

import (
	"errors"
	"testing"
	"time"

	"school_schedule_2/internal/domain/model"

	"github.com/stretchr/testify/assert"
)

func TestUseCase_UpdateClass(t *testing.T) {
	t.Parallel()

	now := time.Now()
	errTest := errors.New("update test error")

	type args struct {
		req UpdateClassReq
	}

	tests := []struct {
		name    string
		args    args
		want    *model.Class
		wantErr error
		before  func(m mockService, args args)
	}{
		{
			name: "success",
			args: args{
				req: UpdateClassReq{
					ID:    1,
					Grade: "9А",
				},
			},
			want: &model.Class{
				ID:        1,
				Grade:     "9А",
				CreatedAt: now,
				UpdatedAt: now,
			},
			before: func(m mockService, args args) {
				class := &model.Class{

					ID:        1,
					Grade:     "9А",
					CreatedAt: now,
					UpdatedAt: now,
				}
				m.classRepo.EXPECT().GetByID(args.req.ID).Return(class, nil)
				updatedClass := *class
				updatedClass.Grade = args.req.Grade

				m.classRepo.EXPECT().Update(&updatedClass).Return(class, nil)
			},
		},
		{
			name: "error on GetByID classes",
			args: args{
				req: UpdateClassReq{
					ID:    1,
					Grade: "9А",
				},
			},
			wantErr: errTest,
			before: func(m mockService, args args) {

				m.classRepo.EXPECT().GetByID(args.req.ID).Return(nil, errTest)
			},
		},
		{
			name: "error on UpdateClass",
			args: args{
				req: UpdateClassReq{
					ID:    1,
					Grade: "9А",
				},
			},
			wantErr: errTest,
			before: func(m mockService, args args) {
				class := &model.Class{

					ID:        1,
					Grade:     "9А",
					CreatedAt: now,
					UpdatedAt: now,
				}
				m.classRepo.EXPECT().GetByID(args.req.ID).Return(class, nil)

				m.classRepo.EXPECT().Update(class).Return(nil, errTest)
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

			classes, err := usecase.Update(tc.args.req)
			if tc.wantErr != nil {
				a.Error(err)
				a.Contains(err.Error(), tc.wantErr.Error())

				return
			}
			a.Equal(tc.want, classes)
			a.NoError(err)
		})
	}
}
