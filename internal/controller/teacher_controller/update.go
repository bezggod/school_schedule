package teacher_controller

import (
	"encoding/json"
	"net/http"
	"school_schedule_2/internal/controller"
	"school_schedule_2/internal/usecase/teacher_usecase"
	"strconv"
)

type updateTeacherReq struct {
	Surname    string `json:"surname"`
	Name       string `json:"name"`
	Patronymic string `json:"patronymic"`
}

func (c *Controller) Update(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")
	teacherID, err := strconv.ParseInt(id, 0, 64)
	if err != nil {
		controller.ValidationErrorRespond(w, controller.NewValidationError("class id not present", "id"))
		return
	}

	// создание структуры из запроса
	decoder := json.NewDecoder(r.Body)
	var req updateTeacherReq
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

	teacher, err := c.teacherUseCase.Update(teacher_usecase.UpdateTeacherReq{
		ID:         teacherID,
		Surname:    req.Surname,
		Name:       req.Name,
		Patronymic: req.Patronymic,
	})
	if err != nil {
		controller.InternalServerErrorRespond(w, err)

		return
	}
	if err = controller.EncodeResponse(w, teacher); err != nil {
		return
	}
}

func validateUpdateRequest(req updateTeacherReq) *controller.ValidationError {
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
