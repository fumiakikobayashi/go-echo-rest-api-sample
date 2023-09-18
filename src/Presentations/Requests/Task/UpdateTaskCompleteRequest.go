package Requests

import "github.com/go-playground/validator/v10"

type UpdateTaskCompleteRequest struct {
	TaskId int `param:"taskId"`
}

func (r *UpdateTaskCompleteRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
