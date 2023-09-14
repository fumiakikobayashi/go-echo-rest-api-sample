package UseCases

import (
	"go-echo-rest-api-sample/src/DomainServices"
	"go-echo-rest-api-sample/src/Domains/Task"
	"go-echo-rest-api-sample/src/Presentations/Requests/Task"
)

type UpdateTaskFavoriteUseCase struct {
	taskRepository DomainServices.ITaskRepository
}

func NewFavoriteTaskUseCase(taskRepository DomainServices.ITaskRepository) *UpdateTaskFavoriteUseCase {
	return &UpdateTaskFavoriteUseCase{
		taskRepository: taskRepository,
	}
}

func (u *UpdateTaskFavoriteUseCase) Execute(request Requests.UpdateTaskFavoriteRequest) error {
	taskId, err := Domains.NewTaskId(request.TaskId)
	if err != nil {
		return err
	}

	task, err := u.taskRepository.GetTask(taskId)
	if err != nil {
		return err
	}

	// お気に入り状態を更新
	task.UpdateTaskFavorite()
	if err := u.taskRepository.UpdateTask(task); err != nil {
		return err
	}
	return nil
}
