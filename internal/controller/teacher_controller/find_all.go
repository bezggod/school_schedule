package teacher_controller

import (
	"encoding/json"
	"net/http"
	"school_schedule_2/internal/controller"
	"school_schedule_2/internal/usecase/teacher_usecase"
)

type findAllTeacherReq struct {
	Surname string `json:"surname"`
	Name    string `json:"name"`
}

func (c *Controller) FindAll(w http.ResponseWriter, r *http.Request) {

	// создание структуры из запроса
	decoder := json.NewDecoder(r.Body)
	var req findAllTeacherReq
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

	teacher, err := c.teacherUseCase.FindAll(teacher_usecase.FindAllTeacherReq{
		Surname: req.Surname,
		Name:    req.Name,
	})
	if err != nil {
		controller.InternalServerErrorRespond(w, err)

		return
	}
	if err = controller.EncodeResponse(w, teacher); err != nil {
		return
	}
}

func validateFindAllRequest(req findAllTeacherReq) *controller.ValidationError {
	if req.Surname == "" {
		return controller.NewValidationError("surname is required", "surname")
	}
	if req.Name == "" {
		return controller.NewValidationError("name is required", "name")
	}

	return nil
}
