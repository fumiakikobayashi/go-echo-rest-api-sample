package UseCases

import (
	"go-ddd-rest-api-sample/sdk"
	Domains "go-ddd-rest-api-sample/src/Domains/Task"
	Requests "go-ddd-rest-api-sample/src/Presentations/Requests/Task"
)

type UpdateTaskFavoriteUseCase struct {
	taskRepository TaskRepositoryInterface
	logger         sdk.LoggerInterface
}

func NewFavoriteTaskUseCase(taskRepository TaskRepositoryInterface, logger sdk.LoggerInterface) *UpdateTaskFavoriteUseCase {
	return &UpdateTaskFavoriteUseCase{
		taskRepository: taskRepository,
		logger:         logger,
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
