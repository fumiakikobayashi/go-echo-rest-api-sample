package Requests

import (
	"github.com/go-playground/validator/v10"
)

type GetTasksRequest struct {
	Sort  string `query:"sort" validate:"omitempty,oneof=name deadline favorite"`
	Order string `query:"order" validate:"omitempty,oneof=asc desc"`
}

func (r *GetTasksRequest) Validate() error {
	return validator.New().Struct(r)
}
