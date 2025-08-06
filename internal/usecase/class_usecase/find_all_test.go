package class_usecase

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
		want    []*model.Class
		wantErr error
		before  func(m mockService, args args)
	}{
		{
			name: "success",
			args: args{
				req: FindAllReq{
					ClassID: 1,
				},
			},
			want: []*model.Class{
				{
					ID:        0,
					Grade:     "",
					CreatedAt: now,
					UpdatedAt: now,
				},
			},
			before: func(m mockService, args args) {
				classes := []*model.Class{
					{
						ID:        0,
						Grade:     "",
						CreatedAt: now,
						UpdatedAt: now,
					},
				}
				m.classRepo.EXPECT().FindAll(dto.FindAllClassesFilter{
					ClassID: args.req.ClassID,
					Grade:   args.req.Grade,
				}).Return(classes, nil)
			},
		},
		{
			name: "classRepo.FindAll err",
			args: args{
				req: FindAllReq{
					ClassID: 1,
				},
			},
			wantErr: errTest,
			before: func(m mockService, args args) {
				m.classRepo.EXPECT().FindAll(dto.FindAllClassesFilter{
					ClassID: args.req.ClassID,
					Grade:   args.req.Grade,
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

			classes, err := usecase.FindAll(tc.args.req)
			if tc.wantErr != nil {
				a.Error(err)
				a.Contains(err.Error(), tc.wantErr.Error())

				return
			}
			a.NoError(err)
			a.Equal(tc.want, classes)
		})
	}
}
