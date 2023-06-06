package UseCases

import Domains "go-ddd-rest-api-sample/src/Domains/Task"

type TaskRepositoryInterface interface {
	GetTasks() (*Domains.TaskList, error)
	GetTask(taskId Domains.TaskId) (*Domains.Task, error)
	SaveTask(task *Domains.Task) error
	UpdateTask(task *Domains.Task) error
}
