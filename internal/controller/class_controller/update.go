package class_controller

import (
	"encoding/json"
	"net/http"
	"school_schedule_2/internal/controller"
	"school_schedule_2/internal/usecase/class_usecase"
	"strconv"
)

type updateClassReq struct {
	Grade string `json:"grade"`
}

func (c *Controller) Update(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")
	classID, err := strconv.ParseInt(id, 0, 64)
	if err != nil {
		controller.ValidationErrorRespond(w, controller.NewValidationError("class id not present", "id"))
		return
	}

	// создание структуры из запроса
	decoder := json.NewDecoder(r.Body)
	var req updateClassReq
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

	class, err := c.classUseCase.Update(class_usecase.UpdateClassReq{
		ID:    classID,
		Grade: req.Grade,
	})
	if err != nil {
		controller.InternalServerErrorRespond(w, err)

		return
	}
	if err = controller.EncodeResponse(w, class); err != nil {
		return
	}
}

func validateUpdateRequest(req updateClassReq) *controller.ValidationError {
	if req.Grade == "" {
		return controller.NewValidationError("name is required", "name")
	}

	return nil
}
