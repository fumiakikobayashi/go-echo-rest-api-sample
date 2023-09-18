package Requests

import (
	"github.com/go-playground/validator/v10"
)

type GetTasksRequest struct {
	Sort  string `query:"sort" validate:"required,oneof=name deadline favorite"`
	Order string `query:"order" validate:"required,oneof=asc desc"`
}

func (r *GetTasksRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
