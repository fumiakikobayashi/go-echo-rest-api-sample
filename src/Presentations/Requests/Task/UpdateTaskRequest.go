package Requests

import "github.com/go-playground/validator/v10"

type UpdateTaskRequest struct {
	TaskId   int    `param:"taskId"`
	Name     string `json:"name"`
	Deadline string `json:"deadline"`
}

func (r *UpdateTaskRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
