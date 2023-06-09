package Requests

type SaveTaskRequest struct {
	Name     string `json:"name"`
	Deadline string `json:"deadline"`
}
