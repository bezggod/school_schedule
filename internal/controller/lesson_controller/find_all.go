package lesson_controller

import (
	"encoding/json"
	"net/http"
	"school_schedule_2/internal/controller"
	"school_schedule_2/internal/domain/model/enums"
	"school_schedule_2/internal/usecase/lesson_usecase"
)

type findAllLessonReq struct {
	TeacherID int64              `json:"teacher_id"`
	ClassID   int64              `json:"class_id"`
	Office    enums.OfficeName   `json:"office"`
	TimeSlot  enums.TimeSlotName `json:"time_slot"`
	Subject   enums.SubjectName  `json:"subject"`
}

func (c *Controller) FindAll(w http.ResponseWriter, r *http.Request) {

	// создание структуры из запроса
	decoder := json.NewDecoder(r.Body)
	var req findAllLessonReq
	err := decoder.Decode(&req)
	if err != nil {
		controller.InternalServerErrorRespond(w, err)

		return
	}

	// валидируем поля запроса
	if validationError := validateFindAllRequest(req); validationError != nil {
		controller.ValidationErrorRespond(w, validationError)

		return
	}

	lesson, err := c.lessonUseCase.FindAll(lesson_usecase.FindAllLessonReq{
		TeacherID: req.TeacherID,
		ClassID:   req.ClassID,
		Office:    req.Office,
		TimeSlot:  req.TimeSlot,
		Subject:   req.Subject,
	})
	if err != nil {
		controller.InternalServerErrorRespond(w, err)

		return
	}
	if err = controller.EncodeResponse(w, lesson); err != nil {
		return
	}
}

func validateFindAllRequest(req findAllLessonReq) *controller.ValidationError {
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

	return nil
}
