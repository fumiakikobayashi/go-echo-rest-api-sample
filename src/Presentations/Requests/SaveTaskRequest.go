package Requests

type SaveTaskRequest struct {
	TaskId   int    `param:"taskId"`
	Name     string `json:"name"`
	Deadline string `json:"deadline"`
}
