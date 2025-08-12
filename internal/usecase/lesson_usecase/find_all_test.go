package lesson_usecase

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"school_schedule_2/internal/domain/dto"
	"school_schedule_2/internal/domain/model"
	"school_schedule_2/internal/domain/model/enums"
	"testing"
	"time"
)

func TestUseCase_FindAll(t *testing.T) {
	t.Parallel()

	now := time.Now()
	errTest := errors.New("test error")

	teacher := &model.Teacher{ID: 1}
	class := &model.Class{ID: 1}

	type args struct {
		req FindAllReq
	}

	tests := []struct {
		name    string
		args    args
		want    *FindAllResponse
		wantErr error
		before  func(m mockService, args args)
	}{
		{
			name: "success",
			args: args{
				req: FindAllReq{
					TeacherID: 1,
					ClassID:   1,
					Office:    enums.Office10,
					TimeSlot:  enums.FirstLesson,
				},
			},
			want: &FindAllResponse{
				Lessons: []*Lesson{
					{
						ID:        1,
						Teacher:   teacher,
						Class:     class,
						Office:    enums.Office10,
						TimeSlot:  enums.FirstLesson,
						Date:      now,
						CreatedAt: now,
						UpdatedAt: now,
					},
				},
			},
			before: func(m mockService, args args) {
				lessons := []*model.Lesson{
					{
						ID:        1,
						TeacherID: teacher.ID,
						ClassID:   class.ID,
						Office:    args.req.Office,
						TimeSlot:  args.req.TimeSlot,
						Subject:   args.req.Subject,
						Date:      now,
						CreatedAt: now,
						UpdatedAt: now,
					},
				}
				m.lessonRepo.EXPECT().FindAll(dto.FindAllLessonFilter{
					TeacherID:  args.req.TeacherID,
					ClassID:    args.req.ClassID,
					OfficeName: args.req.Office,
					Subject:    args.req.Subject,
					TimeSlot:   args.req.TimeSlot,
				}).Return(lessons, nil)

				m.teacherRepo.EXPECT().GetByID(args.req.TeacherID).Return(teacher, nil)
				m.classRepo.EXPECT().GetByID(args.req.ClassID).Return(class, nil)

			},
		},
		{
			name: "lessonRepo.FindAll error",
			args: args{
				req: FindAllReq{
					TeacherID: 1,
					ClassID:   1,
					Office:    enums.Office10,
					TimeSlot:  enums.FirstLesson,
				},
			},
			wantErr: errTest,
			before: func(m mockService, args args) {
				m.lessonRepo.EXPECT().FindAll(dto.FindAllLessonFilter{
					TeacherID:  args.req.TeacherID,
					ClassID:    args.req.ClassID,
					OfficeName: args.req.Office,
					Subject:    args.req.Subject,
					TimeSlot:   args.req.TimeSlot,
				}).Return(nil, errTest)
			},
		},
		{
			name: "teachererRepo.GetByID error",
			args: args{
				req: FindAllReq{
					TeacherID: 1,
					ClassID:   2,
				},
			},
			wantErr: errTest,
			before: func(m mockService, args args) {
				m.lessonRepo.EXPECT().FindAll(dto.FindAllLessonFilter{
					TeacherID:  args.req.TeacherID,
					ClassID:    args.req.ClassID,
					OfficeName: args.req.Office,
					Subject:    args.req.Subject,
					TimeSlot:   args.req.TimeSlot,
				}).Return([]*model.Lesson{
					{
						ID:        1,
						TeacherID: args.req.TeacherID,
						ClassID:   args.req.ClassID,
						Office:    args.req.Office,
						TimeSlot:  args.req.TimeSlot,
						Subject:   args.req.Subject,
						Date:      now,
						CreatedAt: now,
						UpdatedAt: now,
					},
				}, nil)
				m.teacherRepo.EXPECT().GetByID(args.req.TeacherID).Return(nil, errTest)
			},
		},
		{
			name: "teacherRepo.GetByID returrns nil teacher",
			args: args{
				req: FindAllReq{
					TeacherID: 1,
					ClassID:   2,
				},
			},
			wantErr: fmt.Errorf("teacherRepo.GetByID: %w with id %d", errEmptyTeacher, 1),
			before: func(m mockService, args args) {
				m.lessonRepo.EXPECT().FindAll(dto.FindAllLessonFilter{
					TeacherID:  args.req.TeacherID,
					ClassID:    args.req.ClassID,
					OfficeName: args.req.Office,
					Subject:    args.req.Subject,
					TimeSlot:   args.req.TimeSlot,
				}).Return([]*model.Lesson{
					{
						ID:        1,
						TeacherID: args.req.TeacherID,
						ClassID:   args.req.ClassID,
						Office:    args.req.Office,
						TimeSlot:  args.req.TimeSlot,
						Subject:   args.req.Subject,
						Date:      now,
						CreatedAt: now,
						UpdatedAt: now,
					},
				}, nil)
				m.teacherRepo.EXPECT().GetByID(args.req.TeacherID).Return(nil, nil)
			},
		},
		{
			name: "classRepo.GetByID error",
			args: args{
				req: FindAllReq{
					TeacherID: 1,
					ClassID:   2,
				},
			},
			wantErr: errTest,
			before: func(m mockService, args args) {
				m.lessonRepo.EXPECT().FindAll(dto.FindAllLessonFilter{
					TeacherID:  args.req.TeacherID,
					ClassID:    args.req.ClassID,
					OfficeName: args.req.Office,
					Subject:    args.req.Subject,
					TimeSlot:   args.req.TimeSlot,
				}).Return([]*model.Lesson{
					{
						ID:        1,
						TeacherID: args.req.TeacherID,
						ClassID:   args.req.ClassID,
						Office:    args.req.Office,
						TimeSlot:  args.req.TimeSlot,
						Subject:   args.req.Subject,
						Date:      now,
						CreatedAt: now,
						UpdatedAt: now,
					},
				}, nil)
				m.teacherRepo.EXPECT().GetByID(args.req.TeacherID).Return(&model.Teacher{ID: args.req.TeacherID}, nil)
				m.classRepo.EXPECT().GetByID(args.req.ClassID).Return(nil, errTest)
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

			lesson, err := usecase.FindAll(tc.args.req)

			if tc.wantErr != nil {
				a.Error(err)
				a.Contains(err.Error(), tc.wantErr.Error())
				return
			}

			a.NoError(err)
			a.Equal(tc.want, lesson)
		})
	}
}
