package Requests

import "github.com/go-playground/validator/v10"

type DeleteTaskRequest struct {
	TaskId int `param:"taskId" validate:"number,required,min=1"`
}

func (r *DeleteTaskRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
