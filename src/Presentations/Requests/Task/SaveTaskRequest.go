package Requests

import "github.com/go-playground/validator/v10"

type SaveTaskRequest struct {
	Name     string `json:"name"`
	Deadline string `json:"deadline"`
}

func (r *SaveTaskRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
