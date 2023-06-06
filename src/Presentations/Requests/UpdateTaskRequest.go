package Requests

type UpdateTaskRequest struct {
	TaskId   int    `param:"taskId"`
	Name     string `json:"name"`
	Deadline string `json:"deadline"`
}
