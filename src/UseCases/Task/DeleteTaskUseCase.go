package UseCases

import (
	"fmt"
	Domains "go-ddd-rest-api-sample/src/Domains/Task"
	Requests "go-ddd-rest-api-sample/src/Presentations/Requests/Task"
	"go-ddd-rest-api-sample/src/Shared"
)

type DeleteTaskUseCase struct {
	taskRepository TaskRepositoryInterface
	logger         Shared.LoggerInterface
}

func NewDeleteTaskUseCase(taskRepository TaskRepositoryInterface, logger Shared.LoggerInterface) *DeleteTaskUseCase {
	return &DeleteTaskUseCase{
		taskRepository: taskRepository,
		logger:         logger,
	}
}

func (u *DeleteTaskUseCase) Execute(request Requests.DeleteTaskRequest) error {
	taskId, err := Domains.NewTaskId(request.TaskId)
	if err != nil {
		return err
	}

	if err := u.taskRepository.DeleteTask(taskId); err != nil {
		return fmt.Errorf("タスクの削除に失敗しました")
	}
	return nil
}
