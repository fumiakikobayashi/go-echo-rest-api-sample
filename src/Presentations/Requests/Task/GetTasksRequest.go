package Requests

type GetTasksRequest struct {
	Sort  string `query:"sort"`
	Order string `query:"order"`
}
