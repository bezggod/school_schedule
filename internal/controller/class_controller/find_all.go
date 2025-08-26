package class_controller

import (
	"encoding/json"
	"net/http"
	"school_schedule_2/internal/controller"
	"school_schedule_2/internal/usecase/class_usecase"
)

type findAllClassReq struct {
	ClassID int64 `json:"class_id"`
}

func (c *Controller) FindAll(w http.ResponseWriter, r *http.Request) {

	// создание структуры из запроса
	decoder := json.NewDecoder(r.Body)
	var req findAllClassReq
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

	class, err := c.classUseCase.FindAll(class_usecase.FindAllClassReq{
		ClassID: req.ClassID,
	})
	if err != nil {
		controller.InternalServerErrorRespond(w, err)

		return
	}
	if err = controller.EncodeResponse(w, class); err != nil {
		return
	}
}

func validateFindAllRequest(req findAllClassReq) *controller.ValidationError {
	if req.Grade == "" {
		return controller.NewValidationError("name is required", "name")
	}

	return nil
}
