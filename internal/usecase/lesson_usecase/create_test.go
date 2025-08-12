package lesson_usecase

import (
	"errors"
	"testing"
	"time"

	"school_schedule_2/internal/domain/model"
	"school_schedule_2/internal/domain/model/enums"

	"github.com/stretchr/testify/assert"
)

func TestUseCase_CreateLesson(t *testing.T) {
	t.Parallel()

	now := time.Now()
	errTest := errors.New("test error")

	teacher := &model.Teacher{ID: 1}
	class := &model.Class{ID: 1}

	type args struct {
		req CreateLessonReq
	}

	tests := []struct {
		name    string
		args    args
		want    *model.Lesson
		wantErr error
		before  func(m mockService, args args)
	}{
		{
			name: "success",
			args: args{
				req: CreateLessonReq{
					TeacherID: 1,
					ClassID:   1,
					Office:    enums.Office10,
					TimeSlot:  enums.FirstLesson,
					Subject:   enums.SubjectSocialScience,
					Date:      now,
				},
			},
			want: &model.Lesson{
				TeacherID: 1,
				ClassID:   1,
				Office:    enums.Office10,
				TimeSlot:  enums.FirstLesson,
				Subject:   enums.SubjectSocialScience,
				Date:      now,
				CreatedAt: now,
				UpdatedAt: now,
			},
			before: func(m mockService, args args) {
				m.lessonRepo.EXPECT().LessonExists(args.req.Office, args.req.TimeSlot).Return(false)

				m.teacherRepo.EXPECT().GetByID(args.req.TeacherID).Return(teacher, nil)
				m.classRepo.EXPECT().GetByID(args.req.ClassID).Return(class, nil)
				m.timer.EXPECT().NowMoscow().Return(now)

				lesson := &model.Lesson{
					TeacherID: teacher.ID,
					ClassID:   class.ID,
					Office:    args.req.Office,
					TimeSlot:  args.req.TimeSlot,
					Subject:   args.req.Subject,
					Date:      args.req.Date,
					CreatedAt: now,
					UpdatedAt: now,
				}

				m.lessonRepo.EXPECT().Create(lesson).Return(lesson, nil)
			},
		},
		{
			name: "error on create lesson",
			args: args{
				req: CreateLessonReq{
					TeacherID: 1,
					ClassID:   1,
					Office:    enums.Office10,
					TimeSlot:  enums.FirstLesson,
					Subject:   enums.SubjectSocialScience,
					Date:      now,
				},
			},
			wantErr: errTest,
			before: func(m mockService, args args) {
				m.lessonRepo.EXPECT().LessonExists(args.req.Office, args.req.TimeSlot).Return(false)

				m.teacherRepo.EXPECT().GetByID(args.req.TeacherID).Return(teacher, nil)
				m.classRepo.EXPECT().GetByID(args.req.ClassID).Return(class, nil)
				m.timer.EXPECT().NowMoscow().Return(now)

				lesson := &model.Lesson{
					TeacherID: teacher.ID,
					ClassID:   class.ID,
					Office:    args.req.Office,
					TimeSlot:  args.req.TimeSlot,
					Subject:   args.req.Subject,
					Date:      args.req.Date,
					CreatedAt: now,
					UpdatedAt: now,
				}

				m.lessonRepo.EXPECT().Create(lesson).Return(nil, errTest)
			},
		},
		{
			name: "error on get classes by id",
			args: args{
				req: CreateLessonReq{
					TeacherID: 1,
					ClassID:   1,
					Office:    enums.Office10,
					TimeSlot:  enums.FirstLesson,
					Subject:   enums.SubjectSocialScience,
					Date:      now,
				},
			},
			wantErr: errTest,
			before: func(m mockService, args args) {
				m.lessonRepo.EXPECT().LessonExists(args.req.Office, args.req.TimeSlot).Return(false)

				m.teacherRepo.EXPECT().GetByID(args.req.TeacherID).Return(teacher, nil)
				m.classRepo.EXPECT().GetByID(args.req.ClassID).Return(nil, errTest)
			},
		},
		{
			name: "error on get teacher by id",
			args: args{
				req: CreateLessonReq{
					TeacherID: 1,
					ClassID:   1,
					Office:    enums.Office10,
					TimeSlot:  enums.FirstLesson,
					Subject:   enums.SubjectSocialScience,
					Date:      now,
				},
			},
			wantErr: errTest,
			before: func(m mockService, args args) {
				m.lessonRepo.EXPECT().LessonExists(args.req.Office, args.req.TimeSlot).Return(false)

				m.teacherRepo.EXPECT().GetByID(args.req.TeacherID).Return(nil, errTest)
			},
		},
		{
			name: "error lesson exist",
			args: args{
				req: CreateLessonReq{
					TeacherID: 1,
					ClassID:   1,
					Office:    enums.Office10,
					TimeSlot:  enums.FirstLesson,
					Subject:   enums.SubjectSocialScience,
					Date:      now,
				},
			},
			wantErr: errLessonAlreadyExists,
			before: func(m mockService, args args) {
				m.lessonRepo.EXPECT().LessonExists(args.req.Office, args.req.TimeSlot).Return(true)
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

			lesson, err := usecase.CreateLesson(tc.args.req)
			a.ErrorIs(err, tc.wantErr)

			a.Equal(tc.want, lesson)
		})
	}
}
