package UseCases

import (
	Domains "go-ddd-rest-api-sample/src/Domains/Task"
	"go-ddd-rest-api-sample/src/Presentations/Requests"
)

type FavoriteTaskUseCase struct {
	taskRepository TaskRepositoryInterface
}

func NewFavoriteTaskUseCase(taskRepository TaskRepositoryInterface) *FavoriteTaskUseCase {
	return &FavoriteTaskUseCase{
		taskRepository: taskRepository,
	}
}

func (u *FavoriteTaskUseCase) Execute(request Requests.FavoriteTaskRequest) error {
	taskId, err := Domains.NewTaskId(request.TaskId)
	if err != nil {
		return err
	}

	task, err := u.taskRepository.GetTask(taskId)
	if err != nil {
		return err
	}

	// お気に入り状態を更新
	task.FavoriteTask()
	if err := u.taskRepository.UpdateTask(task); err != nil {
		return err
	}
	return nil
}
