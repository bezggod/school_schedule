package class_usecase

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"school_schedule_2/internal/domain/dto"
	"school_schedule_2/internal/domain/model"
	"testing"
	"time"
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
					Grade: nil,
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

					ID: 1,
					Grade: func() string {
						if args.req.Grade != nil {
							return *args.req.Grade
						}
						return "9А"
					}(),
					CreatedAt: now,
					UpdatedAt: now,
				}
				m.classRepo.EXPECT().UpdateClass(dto.UpdateClassFilter{
					ClassID: args.req.ID,
					Grade:   args.req.Grade,
				}).Return(class, nil)
			},
		},
		{
			name: "classRepo.UpdateClass error",
			args: args{
				req: UpdateClassReq{
					ID:    1,
					Grade: nil,
				},
			},
			wantErr: errTest,
			before: func(m mockService, args args) {
				m.classRepo.EXPECT().UpdateClass(dto.UpdateClassFilter{
					ClassID: args.req.ID,
					Grade:   args.req.Grade,
				}).Return(nil, errTest)
			},
		},
		{
			name: "invalid class ID",
			args: args{
				req: UpdateClassReq{
					ID:    0,
					Grade: nil,
				},
			},
			wantErr: fmt.Errorf("invalid class ID"),
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

			classes, err := usecase.UpdateClass(tc.args.req)
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
