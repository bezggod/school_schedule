package class_usecase

import (
	"errors"
	"testing"
	"time"

	"school_schedule_2/internal/domain/model"

	"github.com/stretchr/testify/assert"
)

func TestUseCase_CreateClass(t *testing.T) {
	t.Parallel()

	now := time.Now()
	errTest := errors.New("test error")

	type args struct {
		req CreateClassReq
	}

	tests := []struct {
		name    string
		args    args
		want    *model.Class
		wantErr bool
		before  func(m mockService, args args)
	}{
		{
			name: "success",
			args: args{
				req: CreateClassReq{
					Grade: "7",
				},
			},
			want: &model.Class{
				Grade: "7",
			},
			before: func(m mockService, args args) {
				m.timer.EXPECT().NowMoscow().Return(now)

				class := &model.Class{
					Grade:     args.req.Grade,
					CreatedAt: now,
					UpdatedAt: now,
				}

				m.classRepo.EXPECT().CreateClass(class).Return(&model.Class{
					Grade: "7",
				}, nil)
			},
		},
		{
			name: "error on CreateClass",
			args: args{
				req: CreateClassReq{
					Grade: "7",
				},
			},
			wantErr: true,
			before: func(m mockService, args args) {
				m.timer.EXPECT().NowMoscow().Return(now)

				class := &model.Class{
					Grade:     args.req.Grade,
					CreatedAt: now,
					UpdatedAt: now,
				}
				m.classRepo.EXPECT().CreateClass(class).Return(nil, errTest)
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

			class, err := usecase.CreateClass(tc.args.req)
			if tc.wantErr {
				a.Error(err)

				return
			}
			a.NoError(err)
			a.Equal(tc.want, class)
		})
	}
}
