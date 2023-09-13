package UseCases

import (
	"go-echo-rest-api-sample/src/Domains/Task"
	"go-echo-rest-api-sample/src/UseCases/Shared"
)

type TaskRepositoryInterface interface {
	GetTasks(sortType Shared.SortType, sortOrder Shared.SortOrder) (*Domains.TaskList, error)
	GetTask(taskId Domains.TaskId) (*Domains.Task, error)
	SaveTask(task *Domains.Task) error
	UpdateTask(task *Domains.Task) error
	DeleteTask(taskId Domains.TaskId) error
}
