package Requests

import "github.com/go-playground/validator/v10"

type UpdateTaskFavoriteRequest struct {
	TaskId int `param:"taskId" validate:"number,required,min=1"`
}

func (r *UpdateTaskFavoriteRequest) Validate() error {
	return validator.New().Struct(r)
}
