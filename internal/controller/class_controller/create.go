package class_controller

import (
	"encoding/json"
	"net/http"
	"school_schedule_2/internal/controller"
	"school_schedule_2/internal/usecase/class_usecase"
)

type createClassRequest struct {
	Grade string `json:"grade"`
}

func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {

	// создание структуры из запроса
	decoder := json.NewDecoder(r.Body)
	var req createClassRequest
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

	class, err := c.classUseCase.Create(class_usecase.CreateClassReq{
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

func validateCreateRequest(req createClassRequest) *controller.ValidationError {
	if req.Grade == "" {
		return controller.NewValidationError("name is required", "name")
	}

	return nil
}
