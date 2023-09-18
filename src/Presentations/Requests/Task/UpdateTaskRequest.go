package Requests

import (
	"github.com/go-playground/validator/v10"
	"go-echo-rest-api-sample/src/Presentations/Requests/Shared"
)

type UpdateTaskRequest struct {
	TaskId   int    `param:"taskId" validate:"number,required,min=1"`
	Name     string `json:"name" validate:"number,required,min=1"`
	Deadline string `json:"deadline" validate:"required, dateValidation"`
}

func (r *UpdateTaskRequest) Validate() error {
	validate := validator.New()
	if err := validate.RegisterValidation("dateValidation", Shared.DateValidation); err != nil {
		return err
	}
	return validate.Struct(r)
}
