package class_controller

import (
	"errors"
	"net/http"
	"school_schedule_2/internal/controller"
	"strconv"
	"time"
)

type getClassByIDResp struct {
	ID    int    `json:"id"`
	Grade string `json:"grade"`
}

func (c *Controller) GetByID(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString, 10, 64)
	if err != nil {
		controller.ValidationErrorRespond(w, controller.NewValidationError("id must be integer", "id"))

		return
	}

	class, err := c.classUseCase.GetByID(id)
	if err != nil {
		if errors.Is(err, class_use).NewNotFoundError(w, controller.NewStatusBadRequestError("failed to get client"))

		return
	}

	if err = controller.EncodeResponse(w, mapClientToResponseForGetClientByID(client)); err != nil {
		return
	}
}
