package lesson_controller

import (
	"encoding/json"
	"net/http"
	"school_schedule_2/internal/controller"
	"school_schedule_2/internal/domain/model/enums"
	"school_schedule_2/internal/usecase/lesson_usecase"
	"strconv"
)

type updateLessonReq struct {
	TeacherID int64              `json:"teacher_id"`
	ClassID   int64              `json:"class_id"`
	Office    enums.OfficeName   `json:"office"`
	TimeSlot  enums.TimeSlotName `json:"time_slot"`
	Subject   enums.SubjectName  `json:"subject"`
}

func (c *Controller) Update(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")
	lessonID, err := strconv.ParseInt(id, 0, 64)
	if err != nil {
		controller.ValidationErrorRespond(w, controller.NewValidationError("class id not present", "id"))
		return
	}

	// создание структуры из запроса
	decoder := json.NewDecoder(r.Body)
	var req updateLessonReq
	err = decoder.Decode(&req)
	if err != nil {
		controller.InternalServerErrorRespond(w, err)

		return
	}

	// валидируем поля запроса
	if validationError := validateUpdateRequest(req); validationError != nil {
		controller.ValidationErrorRespond(w, validationError)

		return
	}

	lesson, err := c.lessonUseCase.Update(lesson_usecase.UpdateLessonReq{})
	if err != nil {
		controller.InternalServerErrorRespond(w, err)

		return
	}
	if err = controller.EncodeResponse(w, lesson); err != nil {
		return
	}
}

func validateUpdateRequest(req updateLessonReq) *controller.ValidationError {
	if req.TeacherID == 0 {
		return controller.NewValidationError("teacher_id required", "teacher_id")
	}
	if req.ClassID == 0 {
		return controller.NewValidationError("class_id required", "class_id")
	}
	if req.Office == "" {
		return controller.NewValidationError("office required", "office")
	}
	if req.TimeSlot == 0 {
		return controller.NewValidationError("time_slot required", "time_slot")
	}
	if req.Subject == "" {
		return controller.NewValidationError("subject required", "subject")
	}
	return nil
}
