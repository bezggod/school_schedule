package lesson_controller

import (
	"encoding/json"
	"net/http"
	"school_schedule_2/internal/controller"
	"school_schedule_2/internal/domain/model/enums"
	"school_schedule_2/internal/usecase/lesson_usecase"
	"time"
)

type createLessonReq struct {
	TeacherID int64              `json:"teacher_id"`
	ClassID   int64              `json:"class_id"`
	Office    enums.OfficeName   `json:"office"`
	TimeSlot  enums.TimeSlotName `json:"time_slot"`
	Subject   enums.SubjectName  `json:"subject"`
	Date      time.Time          `json:"date"`
}

func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {

	// создание структуры из запроса
	decoder := json.NewDecoder(r.Body)
	var req createLessonReq
	err := decoder.Decode(&req)
	if err != nil {
		controller.InternalServerErrorRespond(w, err)

		return
	}

	// валидируем поля запроса
	if validationError := validateCreateRequest(req); validationError != nil {
		controller.ValidationErrorRespond(w, validationError)

		return
	}

	lesson, err := c.lessonUseCase.Create(lesson_usecase.CreateLessonReq{
		TeacherID: req.TeacherID,
		ClassID:   req.ClassID,
		Office:    req.Office,
		TimeSlot:  req.TimeSlot,
		Subject:   req.Subject,
		Date:      req.Date,
	})
	if err != nil {
		controller.InternalServerErrorRespond(w, err)

		return
	}
	if err = controller.EncodeResponse(w, lesson); err != nil {
		return
	}
}

func validateCreateRequest(req createLessonReq) *controller.ValidationError {
	if req.TeacherID == 0 {
		return controller.NewValidationError("teacher_id is required", "teacher_id")
	}
	if req.ClassID == 0 {
		return controller.NewValidationError("class_id is required", "class_id")
	}
	if req.Office == "" {
		return controller.NewValidationError("office is required", "office")
	}
	if req.TimeSlot == 0 {
		return controller.NewValidationError("time_slot is required", "time_slot")
	}
	if req.Subject == "" {
		return controller.NewValidationError("subject is required", "subject")
	}
	if req.Date.IsZero() {
		return controller.NewValidationError("date is required", "date")
	}

	return nil
}
