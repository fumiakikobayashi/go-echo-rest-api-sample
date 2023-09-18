package Requests

import "github.com/go-playground/validator/v10"

type GetTaskRequest struct {
	TaskId int `param:"taskId" validate:"number,required,min=1"`
}

func (r *GetTaskRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
