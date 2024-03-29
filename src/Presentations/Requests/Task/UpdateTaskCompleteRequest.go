package Requests

import "github.com/go-playground/validator/v10"

type UpdateTaskCompleteRequest struct {
	TaskId int `param:"taskId" validate:"number,required,min=1"`
}

func (r *UpdateTaskCompleteRequest) Validate() error {
	return validator.New().Struct(r)
}
