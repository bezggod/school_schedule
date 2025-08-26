package teacher_controller

import (
	"encoding/json"
	"net/http"
	"school_schedule_2/internal/controller"
	"school_schedule_2/internal/usecase/teacher_usecase"
)

type createTeacherReq struct {
	Surname    string `json:"surname"`
	Name       string `json:"name"`
	Patronymic string `json:"patronymic"`
}

func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {

	// создание структуры из запроса
	decoder := json.NewDecoder(r.Body)
	var req createTeacherReq
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

	class, err := c.teacherUseCase.Create(teacher_usecase.CreateTeacherReq{
		Surname:    req.Surname,
		Name:       req.Name,
		Patronymic: req.Patronymic,
	})
	if err != nil {
		controller.InternalServerErrorRespond(w, err)

		return
	}
	if err = controller.EncodeResponse(w, class); err != nil {
		return
	}
}

func validateCreateRequest(req createTeacherReq) *controller.ValidationError {
	if req.Surname == "" {
		return controller.NewValidationError("surname is required", "surname")
	}
	if req.Name == "" {
		return controller.NewValidationError("name is required", "name")
	}
	if req.Patronymic == "" {
		return controller.NewValidationError("patronymic is required", "patronymic")
	}

	return nil
}
